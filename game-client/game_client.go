package gameclient

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	game "github.com/h-abranches-dev/connect-4/common"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
	"time"
)

const (
	verifyCompatibilityTimeout = 500 * time.Millisecond
	loginTimeout               = 5 * time.Second

	polTimeout                 = 500 * time.Millisecond
	polInterval                = 2 * time.Second
	checkMatchCanStartTimeout  = 500 * time.Millisecond
	checkMatchCanStartInterval = 2 * time.Second

	serveBoardTimeout = 500 * time.Millisecond
	playTimeout       = 10000 * time.Millisecond
)

var (
	gsAddr  string
	version = new(versions.Version)

	rc gameserver.RouteClient
)

func SetGSAddr(host string, port int) {
	gsAddr = utils.NewAddress(host, port)
}

func GetGSAddr() string {
	return gsAddr
}

func SetVersion(v string) error {
	pv, err := versions.GetVersion(versions.Get(), v)
	if err != nil {
		return err
	}
	*version = *pv
	return nil
}

func DisplayVersion() {
	v := colors.FgRed(version.Tag)
	fmt.Printf("version: %s\n\n", v)
}

// DisplayGSAddr optional
func DisplayGSAddr(addr string) {
	fmt.Printf("game server address: %s\n\n", addr)
}

// OpenNewConnWithGameServer set up a connection to the game server creating a route client
func OpenNewConnWithGameServer(addr string) (*grpc.ClientConn, gameserver.RouteClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// to enter here you can not pass the credentials, example: conn, err := grpc.Dial(*geAddr)
	if err != nil {
		return nil, nil, fmt.Errorf(game.ConnNotPossibleErr, addr)
	}

	return conn, gameserver.NewRouteClient(conn), nil
}

func SetGSRouteClient(nrc gameserver.RouteClient) {
	rc = nrc
}

func VerifyCompatibility() error {
	gcv := version.Tag

	ctx, cancel := context.WithTimeout(context.Background(), verifyCompatibilityTimeout)
	defer cancel()

	if _, err := rc.VerifyCompatibility(ctx, &gameserver.VerifyCompatibilityPayload{
		GameClientVersion: gcv,
	}); err != nil {
		return err
	}

	return nil
}

func Login() (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), loginTimeout)
	defer cancel()

	loginResp, err := rc.Login(ctx, &gameserver.LoginPayload{})
	if err != nil {
		return nil, err
	}

	sessionToken, err := uuid.Parse(loginResp.SessionToken)
	if err != nil {
		return nil, fmt.Errorf(game.InvalidSessionTokenFormatErr)
	}

	return &sessionToken, nil
}

func SendPOL(payload *gameserver.POLPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), polTimeout)
	defer cancel()

	if _, err := rc.POL(ctx, payload); err != nil {
		return err
	}

	return nil
}

func StartSendingPOL(sessionToken uuid.UUID) error {
	ticker := time.NewTicker(polInterval)

	for {
		select {
		case <-ticker.C:
			if err := SendPOL(&gameserver.POLPayload{
				SessionToken: sessionToken.String(),
			}); err != nil {
				ticker.Stop()
				if strings.Contains(err.Error(), game.InvalidSessionTokenFormatErr) ||
					strings.Contains(err.Error(), game.InvalidSessionTokenErr) {
					return fmt.Errorf("%s => %s", game.CannotConnWithGSErr, err.Error())
				} else {
					return fmt.Errorf("%s => %s", game.NoConnWithGSErr, err.Error())
				}
			}
		}
	}
}

func checkMatchCanStart(payload *gameserver.CheckMatchCanStartPayload, checkMatchCanStartCh chan bool,
	checkMatchCanStartErrCh chan error) {

	ctx, cancel := context.WithTimeout(context.Background(), checkMatchCanStartTimeout)
	defer cancel()

	resp, err := rc.CheckMatchCanStart(ctx, payload)
	if err != nil {
		checkMatchCanStartErrCh <- err
	}

	if resp != nil && resp.CanStart {
		checkMatchCanStartCh <- resp.CanStart
	} else {
		fmt.Printf("\rWaiting for other play to join...")
	}
}

func StartCheckingMatchCanStart(sessionToken uuid.UUID, errCh chan error) (bool, error) {
	ticker := time.NewTicker(checkMatchCanStartInterval)
	checkMatchCanStartCh := make(chan bool)
	checkMatchCanStartErrCh := make(chan error)

	for {
		select {
		case <-checkMatchCanStartCh:
			ticker.Stop()
			return true, nil
		case err := <-checkMatchCanStartErrCh:
			ticker.Stop()
			return false, err
		case err := <-errCh:
			ticker.Stop()
			return false, err
		case <-ticker.C:
			go checkMatchCanStart(&gameserver.CheckMatchCanStartPayload{
				SessionToken: sessionToken.String(),
			}, checkMatchCanStartCh, checkMatchCanStartErrCh)
		}
	}
}

func GetPlayerIDAndBoard(sessionToken uuid.UUID) (*string, *string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), serveBoardTimeout)
	defer cancel()

	resp, err := rc.ServeBoard(ctx, &gameserver.ServeBoardPayload{
		SessionToken: sessionToken.String(),
	})
	if resp == nil && err != nil {
		return nil, nil, err
	}

	return &resp.PlayerID, &resp.Board, err
}

func Play(sessionToken uuid.UUID, column int) (bool, error) {
	fmt.Printf("\n")
	ctx, cancel := context.WithTimeout(context.Background(), playTimeout)
	defer cancel()
	resp, err := rc.Play(ctx, &gameserver.PlayPayload{
		SessionToken: sessionToken.String(),
		Column:       int32(column),
	})
	if err != nil {
		return false, err
	}

	if resp.ThereIsWinner {
		if resp.Error != "" {
			return true, fmt.Errorf(resp.Error)
		} else {
			return true, nil
		}
	}

	return false, nil
}
