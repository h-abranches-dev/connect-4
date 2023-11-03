package gameclient

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-abranches-dev/connect-4/common"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type sendPOLService struct {
	timeout      time.Duration
	rc           gameserver.RouteClient
	sessionToken uuid.UUID
}

const (
	verifyCompatibilityTimeout = 500 * time.Millisecond
	loginTimeout               = 5 * time.Second

	polTimeout  = 500 * time.Millisecond
	polInterval = 2 * time.Second
)

var (
	version = new(versions.Version)
)

func SetGSAddr(addr *string, host string, port int) {
	*addr = utils.NewAddress(host, port)
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

func DisplayGSAddr(addr string) {
	fmt.Printf("game server address: %s\n\n", addr)
}

// OpenNewConn set up a connection to the game server creating a route client
func OpenNewConn(gsAddr string) (*grpc.ClientConn, gameserver.RouteClient, error) {
	conn, err := grpc.Dial(gsAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("it wasn't possible to create a connection with the game server in %s\n", gsAddr)
	}

	return conn, gameserver.NewRouteClient(conn), nil
}

func VerifyCompatibility(rc gameserver.RouteClient) error {
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

func Login(rc gameserver.RouteClient) (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), loginTimeout)
	defer cancel()

	loginResp, err := rc.Login(ctx, &gameserver.LoginPayload{})
	if err != nil {
		return nil, err
	}

	sessionToken, err := uuid.Parse(loginResp.Token)
	if err != nil {
		return nil, err
	}

	return &sessionToken, nil
}

func newSendPOLService(rc gameserver.RouteClient, timeout time.Duration, sessionToken uuid.UUID) *sendPOLService {
	return &sendPOLService{
		timeout:      timeout,
		rc:           rc,
		sessionToken: sessionToken,
	}
}

func errWhenPOLfails() {
	fmt.Printf("can't talk with game server\n")
	fmt.Printf("game is quitting...\n")
	fmt.Printf("i'm sorry for the inconvenience, bye!\n\n")
}

func (s sendPOLService) Run(clientWillTerminate chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	resp, err := s.rc.POL(ctx, &gameserver.POLPayload{
		SessionToken: s.sessionToken.String(),
	})
	if err != nil {
		errWhenPOLfails()
		clientWillTerminate <- true
		return
	}
	if resp.Err != "" {
		errWhenPOLfails()
		clientWillTerminate <- true
		return
	}
}

func SendPOL(rc gameserver.RouteClient, sessionToken uuid.UUID, stopGame chan bool) {
	stopPOLTicker := make(chan bool)
	sendPOLServ := newSendPOLService(rc, polTimeout, sessionToken)
	common.StartTicker(polInterval, sendPOLServ, stopPOLTicker)
	select {
	case <-stopPOLTicker:
		stopGame <- true
	}
}
