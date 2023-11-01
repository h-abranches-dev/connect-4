package gameserver

import (
	"context"
	"fmt"
	gameengine "github.com/h-abranches-dev/connect-4/game-engine"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

const (
	pingTimeout = 500 * time.Millisecond
)

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
