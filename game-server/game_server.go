package gameserver

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	game "github.com/h-abranches-dev/connect-4/common"
	gameengine "github.com/h-abranches-dev/connect-4/game-engine"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"strings"
	"time"
)

const (
	verifyCompatibilityTimeout = 500 * time.Millisecond
	connectTimeout             = 500 * time.Millisecond

	polTimeout  = 500 * time.Millisecond
	polInterval = 2 * time.Second
)

var (
	geAddr  string
	version = new(versions.Version)

	rc gameengine.RouteClient
)

func SetGEAddr(host string, port int) {
	geAddr = utils.NewAddress(host, port)
}

func GetGEAddr() string {
	return geAddr
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
	gameTitle := colors.FgRed("CONNECT-4")
	v := colors.FgRed(version.Tag)
	fmt.Printf("\n%s game server\n\n", gameTitle)
	fmt.Printf("version: %s\n\n", v)
}

// OpenNewConnWithGameEngine set up a connection to the game engine creating a route client
func OpenNewConnWithGameEngine(addr string) (*grpc.ClientConn, gameengine.RouteClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// to enter here you can not pass the credentials, example: conn, err := grpc.Dial(*geAddr)
	if err != nil {
		return nil, nil, fmt.Errorf(game.ConnNotPossibleErr, addr)
	}

	return conn, gameengine.NewRouteClient(conn), nil
}

func SetGERouteClient(nrc gameengine.RouteClient) {
	rc = nrc
}

func VerifyCompatibility() error {
	gsv := version.Tag

	ctx, cancel := context.WithTimeout(context.Background(), verifyCompatibilityTimeout)
	defer cancel()

	if _, err := rc.VerifyCompatibility(ctx, &gameengine.VerifyCompatibilityPayload{
		GameServerVersion: gsv,
	}); err != nil {
		return err
	}

	return nil
}

// DisplayGEAddr optional
func DisplayGEAddr() {
	fmt.Printf("game engine address: %s\n\n", geAddr)
}

func Connect() (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	connectResp, err := rc.Connect(ctx, &gameengine.ConnectPayload{})
	if err != nil {
		return nil, err
	}

	sessionToken, err := uuid.Parse(connectResp.SessionToken)
	if err != nil {
		return nil, fmt.Errorf(game.InvalidSessionTokenFormatErr)
	}

	return &sessionToken, nil
}

func encodeBoards() string {
	var boardsIDs []string
	for _, m := range *activeMatches {
		boardsIDs = append(boardsIDs, m.BoardID.String())
	}
	return strings.Join(boardsIDs, ";")
}

func removeOutdatedSession() {
	var sessionsToRemove []string
	for _, os := range *openSessions {
		expirationTime := os.LastPOL.Add(sessionLastPOLExpirationInterval)
		now := time.Now()
		if expirationTime.Before(now) {
			sessionsToRemove = append(sessionsToRemove, os.Token.String())
		}
	}
	if len(sessionsToRemove) > 0 {
		for _, sid := range sessionsToRemove {
			removeGCSession(openSessions, sid)
		}
	}
}

func removeOutdatedMatches() {
	for _, m := range *activeMatches {
		if m.IsFinished() && (m.P1 == nil || m.P2 == nil) {
			RemoveMatch(activeMatches, m.Id)
			break
		}
		if m.StatusCode == Started && (m.P1 == nil || m.P2 == nil) {
			m.StatusCode = Abandoned
		}
	}
}

func SendPOL(payload *gameengine.POLPayload) error {
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
			removeOutdatedSession()
			removeOutdatedMatches()

			if err := SendPOL(&gameengine.POLPayload{
				SessionToken:  sessionToken.String(),
				EncodedBoards: encodeBoards(),
			}); err != nil {
				ticker.Stop()
				if strings.Contains(err.Error(), game.InvalidSessionTokenFormatErr) ||
					strings.Contains(err.Error(), game.InvalidSessionTokenErr) {
					return fmt.Errorf("%s => %s", game.CannotConnWithGEErr, err.Error())
				} else {
					return fmt.Errorf("%s => %s", game.NoConnWithGEErr, err.Error())
				}
			}
		}
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
