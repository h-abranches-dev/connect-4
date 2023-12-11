package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-abranches-dev/connect-4/game-client"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/version-go/ldflags"
	"google.golang.org/grpc"
	"strconv"
	"strings"
	"time"
)

const (
	gsDefaultHost = "127.0.0.1"
	gsDefaultPort = 50052

	checkBoardUpdatedTimeout  = 500 * time.Millisecond
	checkBoardUpdatedInterval = 2 * time.Second
)

var (
	gsHost = flag.String("gsHost", gsDefaultHost, "the game server host")
	gsPort = flag.Int("gsPort", gsDefaultPort, "the game server port")
	gsAddr string

	rc gameserver.RouteClient

	errCh = make(chan error)
)

func main() {
	flag.Parse()

	gameclient.SetGSAddr(*gsHost, *gsPort)

	gsAddr = gameclient.GetGSAddr()

	if err := gameclient.SetVersion(ldflags.Version()); err != nil {
		utils.PrintError(err)
		return
	}

	var conn *grpc.ClientConn
	var err error
	conn, rc, err = gameclient.OpenNewConnWithGameServer(gsAddr)
	if err != nil {
		utils.PrintError(err)
		return
	}
	defer utils.CloseConn(conn)
	gameclient.SetGSRouteClient(rc)

	if err = gameclient.VerifyCompatibility(); err != nil {
		utils.PrintError(err)
		return
	}

	var sessionToken *uuid.UUID
	sessionToken, err = start()
	if err != nil {
		utils.PrintError(err)
		return
	}

	// wait for the match to start or for some error
	matchCanStartCh := make(chan struct{})
	go func() {
		matchCanStart := false
		matchCanStart, err = gameclient.StartCheckingMatchCanStart(*sessionToken, errCh)
		if err != nil || matchCanStart {
			<-matchCanStartCh
		}
	}()
	matchCanStartCh <- struct{}{}

	if err != nil {
		utils.PrintError(err)
		return
	}

	var board *string
	var playerID *string
	playerTurn := "P1"
	turnID := 0
	thereIsWinner := false
	for !thereIsWinner {
		turnID++
		utils.ClearConsole()
		playerID, board, err = gameclient.GetPlayerIDAndBoard(*sessionToken)
		if err != nil {
			utils.PrintError(err)
			return
		}
		renderPlay(turnID, playerTurn, *playerID, *board)

		var playResult string
		if playerTurn == *playerID {
			var col int
			fmt.Printf("Choose the column where you want to play (1 - 7), resign (R), or exit (E): ")
			var input string
			_, err = fmt.Scanln(&input)

			for err != nil || !validPlay(input) {
				utils.ClearConsole()
				renderPlay(turnID, playerTurn, *playerID, *board)
				fmt.Printf("\n")
				fmt.Printf("invalid option\n\n")
				fmt.Printf("Choose the column where you want to play (1 - 7), resign (R), or exit (E): ")
				_, err = fmt.Scanln(&input)
			}
			if strings.ToUpper(input) == "R" {
				thereIsWinner = true
				playResult = "R"
			} else if strings.ToUpper(input) == "E" {
				thereIsWinner = true
				playResult = "E"
			} else {
				col, _ = strconv.Atoi(input)
				thereIsWinner, err = gameclient.Play(*sessionToken, col)
				if err != nil {
					utils.PrintError(err)
					return
				}
			}
		} else {
			fmt.Printf("Wait for your turn...\n")
			playResult, err = checkIfBoardWasUpdated(*sessionToken)
			if err != nil {
				utils.PrintError(err)
			} else {
				if playResult == "W" || playResult == "RA" {
					thereIsWinner = true
				}
			}
		}

		if thereIsWinner {
			playerID, board, err = gameclient.GetPlayerIDAndBoard(*sessionToken)
			if err != nil {
				utils.PrintError(err)
				return
			}

			utils.ClearConsole()
			renderPlay(turnID, playerTurn, *playerID, *board)

			if *playerID == playerTurn {
				if playResult == "R" {
					fmt.Printf("You have resigned, thus you have lost the match!\n")
					fmt.Printf("Thank you for playing! Bye!\n\n")
					break
				}
				if playResult == "E" {
					fmt.Printf("You have exited from the game, thus you have lost the match!\n")
					fmt.Printf("Thank you for playing! Bye!\n\n")
					break
				}
				fmt.Printf("You have won the match!\n")
				fmt.Printf("Thank you for playing! Bye!\n\n")
			} else {
				if playResult == "RA" {
					fmt.Printf("Your opponent resigned or abandoned the game, thus you have won the match!\n")
					fmt.Printf("Thank you for playing! Bye!\n\n")
					break
				}
				fmt.Printf("You have lost the match!\n")
				fmt.Printf("Thank you for playing! Bye!\n\n")
			}
			break
		}

		if playerTurn == "P1" {
			playerTurn = "P2"
		} else {
			playerTurn = "P1"
		}
	}
}

func renderPlay(turnID int, playerTurn, playerID, board string) {
	fmt.Printf("%s\n\n", playerID)
	fmt.Printf("%s TURN\n\n", playerTurn)
	if turnID > 0 {
		fmt.Printf("TURN ID %d\n\n", turnID)
	}
	fmt.Printf("%s\n", board)
}

func validPlay(play string) bool {
	if strings.ToUpper(play) == "R" || strings.ToUpper(play) == "E" {
		return true
	}
	col, err := strconv.Atoi(play)
	if err == nil {
		if col >= 1 && col <= 7 {
			return true
		}
	}
	return false
}

func checkIfBoardWasUpdated(sessionToken uuid.UUID) (string, error) {
	ticker := time.NewTicker(checkBoardUpdatedInterval)
	checkIfBoardWasUpdatedCh := make(chan struct{})
	var update string
	var err error
	for {
		select {
		case <-checkIfBoardWasUpdatedCh:
			ticker.Stop()
			return update, err
		case <-ticker.C:
			go func(sessionToken uuid.UUID, checkIfBoardWasUpdatedCh chan struct{}) {
				ctx, cancel := context.WithTimeout(context.Background(), checkBoardUpdatedTimeout)
				defer cancel()

				var resp *gameserver.CheckBoardUpdatedResponse
				resp, err = rc.CheckBoardUpdated(ctx, &gameserver.CheckBoardUpdatedPayload{
					SessionToken: sessionToken.String(),
				})

				if err != nil {
					checkIfBoardWasUpdatedCh <- struct{}{}
				} else if err == nil && resp != nil && resp.Feedback == "W" || resp.Feedback == "P" ||
					resp.Feedback == "RA" {

					update = resp.Feedback
					checkIfBoardWasUpdatedCh <- struct{}{}
				}
			}(sessionToken, checkIfBoardWasUpdatedCh)
		}
	}
}

func initialize() {
	gameTitle := colors.FgRed("CONNECT-4")
	fmt.Printf("\n%s\n\n", gameTitle)

	gameclient.DisplayVersion()
	//gameclient.DisplayGSAddr(gsAddr)
}

func start() (*uuid.UUID, error) {
	utils.ClearConsole()

	initialize()

	fmt.Printf("Write 'START': ")
	var input string
	_, err := fmt.Scanln(&input)
	for err != nil || strings.ToUpper(input) != "START" {
		utils.ClearConsole()

		fmt.Printf("invalid option\n\n")

		fmt.Printf("Write 'START': ")
		_, err = fmt.Scanln(&input)
	}

	fmt.Println()

	sessionToken, err := gameclient.Login()
	if err != nil {
		return nil, err
	}

	go func() {
		err = gameclient.StartSendingPOL(*sessionToken)
		if err != nil {
			errCh <- err
			return
		}
	}()

	return sessionToken, nil
}
