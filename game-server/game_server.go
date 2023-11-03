package gameserver

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-abranches-dev/connect-4/common"
	gameengine "github.com/h-abranches-dev/connect-4/game-engine"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"time"
)

type sendPOLService struct {
	timeout      time.Duration
	rc           gameengine.RouteClient
	sessionToken uuid.UUID
}

const (
	verifyCompatibilityTimeout = 500 * time.Millisecond
	connectTimeout             = 5 * time.Second

	polTimeout  = 500 * time.Millisecond
	polInterval = 2 * time.Second
)

var (
	version = new(versions.Version)
)

func SetGEAddr(addr *string, host string, port int) {
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

// OpenNewConn set up a connection to the game engine creating a route client
func OpenNewConn(geAddr string) (*grpc.ClientConn, gameengine.RouteClient, error) {
	conn, err := grpc.Dial(geAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("it wasn't possible to create a connection with the game engine in %s\n", geAddr)
	}

	return conn, gameengine.NewRouteClient(conn), nil
}

func displayVersion() {
	gameTitle := colors.FgRed("CONNECT-4")
	v := colors.FgRed(version.Tag)
	fmt.Printf("%s game server\n\n", gameTitle)
	fmt.Printf("version: %s\n\n", v)
}

func VerifyCompatibility(rc gameengine.RouteClient) error {
	gsv := version.Tag

	ctx, cancel := context.WithTimeout(context.Background(), verifyCompatibilityTimeout)
	defer cancel()

	if _, err := rc.VerifyCompatibility(ctx, &gameengine.VerifyCompatibilityPayload{
		GameServerVersion: gsv,
	}); err != nil {
		return err
	}

	displayVersion()

	return nil
}

func DisplayGEAddr(addr string) {
	fmt.Printf("game engine address: %s\n\n", addr)
}

func Connect(rc gameengine.RouteClient) (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	connectResp, err := rc.Connect(ctx, &gameengine.ConnectPayload{})
	if err != nil {
		return nil, err
	}

	sessionToken, err := uuid.Parse(connectResp.Token)
	if err != nil {
		return nil, err
	}

	return &sessionToken, nil
}

func newSendPOLService(rc gameengine.RouteClient, timeout time.Duration, sessionToken uuid.UUID) *sendPOLService {
	return &sendPOLService{
		timeout:      timeout,
		rc:           rc,
		sessionToken: sessionToken,
	}
}

func errWhenPOLfails() {
	log.Printf("can't talk with game engine\n")
	log.Printf("game server is quitting...\n")
	log.Printf("i'm sorry for the inconvenience, bye!\n\n")
}

func (s sendPOLService) Run(clientWillTerminate chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	resp, err := s.rc.POL(ctx, &gameengine.POLPayload{
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

func SendPOL(rc gameengine.RouteClient, sessionToken uuid.UUID, stopGS chan bool) {
	stopPOLTicker := make(chan bool)
	sendPOLServ := newSendPOLService(rc, polTimeout, sessionToken)
	common.StartTicker(polInterval, sendPOLServ, stopPOLTicker)
	select {
	case <-stopPOLTicker:
		stopGS <- true
	}
}

func StartGRPCServer(port int) error {
	listSrvAddr := utils.ListSrvAddr(port)
	lis, err := net.Listen("tcp", listSrvAddr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	RegisterRouteServer(grpcServer, NewGameServer())

	log.Printf(">>> game server is listening on the address %s\n\n", colors.FgRed(lis.Addr().String()))

	if err = grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
