package gameclient

import (
	"context"
	"fmt"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

const (
	pingTimeout = 500 * time.Millisecond
)

// OpenNewConn set up a connection to the game server creating a route client
func OpenNewConn(gsAddr string) (*grpc.ClientConn, gameserver.RouteClient) {
	conn, err := grpc.Dial(gsAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("it wasn't possible to create a connection with the game server in %s\n", gsAddr)
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
