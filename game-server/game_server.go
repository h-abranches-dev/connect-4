package gameserver

import (
	"context"
	"fmt"
	gameengine "github.com/h-abranches-dev/connect-4/game-engine"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type GameServer struct {
	UnimplementedRouteServer
}

const (
	pingTimeout = 500 * time.Millisecond
)

var (
	version = new(versions.Version)
)

func NewGameServer() *GameServer {
	return &GameServer{}
}

func (gs *GameServer) Ping(ctx context.Context, payload *PingPayload) (*PingResponse, error) {
	return &PingResponse{Pong: "pong"}, nil
}

func (gs *GameServer) VerifyCompatibility(ctx context.Context,
	payload *VerifyCompatibilityPayload) (*VerifyCompatibilityResponse, error) {

	gsv := version

	gcv := versions.Version{
		Tag: payload.GameClientVersion,
	}
	if !gsv.Supports(gcv) {
		return nil, utils.FormatErrors([]string{
			fmt.Sprintf("game client version %q is not compatible with the game server version %q.",
				gcv.Tag, gsv.Tag),
			fmt.Sprintf("game client versions compatible with game server version %q: %s", gsv.Tag, versions.GetGameClientsVersions(gsv.SupportedVersions)),
		})
	}

	return &VerifyCompatibilityResponse{GameServerVersion: gsv.Tag}, nil
}

func GetVersion() string {
	return version.Tag
}

func SetVersion(nv versions.Version) {
	*version = nv
}

// OpenNewConn set up a connection to the game engine creating a route client
func OpenNewConn(geAddr string) (*grpc.ClientConn, gameengine.RouteClient) {
	conn, err := grpc.Dial(geAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("it wasn't possible to create a connection with the game engine in %s\n", geAddr)
	}

	return conn, gameengine.NewRouteClient(conn)
}

func PingGameEngine(rc gameengine.RouteClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	resp, err := rc.Ping(ctx, &gameengine.PingPayload{})
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	fmt.Printf("game server: ping\n")
	fmt.Printf("game engine: %s\n\n", resp.Pong)

	return nil
}
