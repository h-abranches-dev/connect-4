package main

import (
	"context"
	"flag"
	"fmt"
	gameengine "github.com/h-abranches-dev/connect-4/game-engine"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"github.com/version-go/ldflags"
	"google.golang.org/grpc"
	"net"
	"time"
)

const (
	defaultPort   = 50052
	geDefaultHost = "127.0.0.1"
	geDefaultPort = 50051

	verifyCompatibilityTimeout = 500 * time.Millisecond
)

var (
	port   = flag.Int("port", defaultPort, "the game server port")
	geHost = flag.String("geHost", geDefaultHost, "the game engine host")
	gePort = flag.Int("gePort", geDefaultPort, "the game engine port")
	geAddr string
)

func main() {
	flag.Parse()

	setGEAddr()
	setVersion(ldflags.Version())

	conn, rc := gameserver.OpenNewConn(geAddr)
	defer utils.CloseConn(conn)

	verifyCompatibility(rc)
	displayGEAddr()

	err := gameserver.PingGameEngine(rc)
	if err != nil {
		fmt.Printf("%s\n\n", err.Error())
		return
	}

	listSrvAddr := utils.ListSrvAddr(*port)
	lis, err := net.Listen("tcp", listSrvAddr)
	if err != nil {
		fmt.Printf("failed to create listener: %s\n\n", err.Error())
		return
	}

	grpcServer := grpc.NewServer()
	gameserver.RegisterRouteServer(grpcServer, gameserver.NewGameServer())

	fmt.Printf("game server is listening on the address %s\n", colors.FgRed(lis.Addr().String()))

	if err = grpcServer.Serve(lis); err != nil {
		fmt.Printf("%s\n\n", err.Error())
		return
	}
}

func setVersion(v string) {
	pv, err := versions.GetVersion(versions.Get(), v)
	if err != nil {
		panic(err)
	}
	gameserver.SetVersion(*pv)
}

func displayVersion() {
	gameTitle := colors.FgRed("CONNECT-4")
	v := colors.FgRed(gameserver.GetVersion())
	fmt.Printf("%s game server\n\n", gameTitle)
	fmt.Printf("version: %s\n\n", v)
}

func setGEAddr() {
	geAddr = utils.NewAddress(*geHost, *gePort)
}

func displayGEAddr() {
	fmt.Printf("game engine address: %s\n\n", geAddr)
}

func verifyCompatibility(rc gameengine.RouteClient) {
	gsv := gameserver.GetVersion()

	ctx, cancel := context.WithTimeout(context.Background(), verifyCompatibilityTimeout)
	defer cancel()

	_, err := rc.VerifyCompatibility(ctx, &gameengine.VerifyCompatibilityPayload{
		GameServerVersion: gsv,
	})
	if err != nil {
		panic(err)
	}

	displayVersion()
}
