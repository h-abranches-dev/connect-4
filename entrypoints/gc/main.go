package main

import (
	"flag"
	"fmt"
	"github.com/h-abranches-dev/connect-4/game-client"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/version-go/ldflags"
	"strings"
)

const (
	gsDefaultHost = "127.0.0.1"
	gsDefaultPort = 50052
)

var (
	gsHost = flag.String("gsHost", gsDefaultHost, "the game server host")
	gsPort = flag.Int("gsPort", gsDefaultPort, "the game server port")
	gsAddr = new(string)

	stopGame = make(chan bool)
)

func main() {
	flag.Parse()

	gameclient.SetGSAddr(gsAddr, *gsHost, *gsPort)

	if err := gameclient.SetVersion(ldflags.Version()); err != nil {
		utils.PrintError(err)
		return
	}

	conn, rc, err := gameclient.OpenNewConn(*gsAddr)
	if err != nil {
		utils.PrintError(err)
		return
	}
	defer utils.CloseConn(conn)

	if err = gameclient.VerifyCompatibility(rc); err != nil {
		utils.PrintError(err)
		return
	}

	if err = start(*gsAddr, rc); err != nil {
		fmt.Printf("\n%s\n", err.Error())
	}
}

func initialize(gsAddr string) {
	gameTitle := colors.FgRed("CONNECT-4")
	fmt.Printf("\n%s\n\n", gameTitle)

	gameclient.DisplayVersion()
	gameclient.DisplayGSAddr(gsAddr)
}

func start(gsAddr string, rc gameserver.RouteClient) error {
	utils.ClearConsole()

	initialize(gsAddr)

	fmt.Printf("Write 'START': ")
	var input string
	_, err := fmt.Scanln(&input)
	for err != nil || strings.ToUpper(input) != "START" {
		utils.ClearConsole()

		fmt.Printf("invalid option\n\n")

		fmt.Printf("Write 'START': ")
		_, err = fmt.Scanln(&input)
	}

	sessionToken, err := gameclient.Login(rc)
	if err != nil {
		utils.PrintError(err)
		return err
	}

	go gameclient.SendPOL(rc, *sessionToken, stopGame)

	fmt.Printf("\nlogin has succeeded\n\n")

	<-stopGame

	return nil
}
