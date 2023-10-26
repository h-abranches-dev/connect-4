package gameclient

import (
	"context"
	"fmt"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

const (
	pingTimeout = 500 * time.Millisecond
)

var (
	gameClientVersion = new(versions.Version)
)

func GetVersion() *versions.Version {
	return gameClientVersion
}

func SetVersion(v *versions.Version, nv versions.Version) {
	*v = nv
}

// OpenNewConn set up a connection to the game server creating a route client
func OpenNewConn(gsAddr string) (*grpc.ClientConn, gameserver.RouteClient) {
	conn, err := grpc.Dial(gsAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("it wasn't possible creating a client connection with the game server in %s\n", gsAddr)
	}

	return conn, gameserver.NewRouteClient(conn)
}

func PingGameServer(rc gameserver.RouteClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	resp, err := rc.Ping(ctx, &gameserver.PingPayload{})
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	fmt.Printf("game client: ping\n")
	fmt.Printf("game server: %s\n\n", resp.Pong)

	return nil
}
