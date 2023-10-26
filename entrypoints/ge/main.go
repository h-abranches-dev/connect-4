package main

import (
	"flag"
	"fmt"
	gameengine "github.com/h-abranches-dev/connect-4/game-engine"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"github.com/version-go/ldflags"
	"google.golang.org/grpc"
	"net"
)

const (
	defaultPort = 50051
)

var (
	port = flag.Int("port", defaultPort, "the game engine port")
)

func main() {
	setVersion(ldflags.Version())

	displayVersion()

	flag.Parse()

	listSrvAddr := utils.ListSrvAddr(*port)
	lis, err := net.Listen("tcp", listSrvAddr)
	if err != nil {
		fmt.Printf("failed to create listener: %s\n\n", err.Error())
		return
	}

	grpcServer := grpc.NewServer()
	gameengine.RegisterRouteServer(grpcServer, gameengine.NewGameEngineWrapper())

	fmt.Printf("game engine is listening on the address %s\n", colors.FgRed(lis.Addr().String()))

	if err = grpcServer.Serve(lis); err != nil {
		fmt.Printf("%s\n\n", err.Error())
		return
	}
}

func setVersion(v string) {
	pv, err := versions.Set(v)
	if err != nil {
		panic(err)
	}
	gameengine.SetVersion(gameengine.GetVersion(), *pv)
}

func displayVersion() {
	gameTitle := colors.FgRed("CONNECT-4")
	v := colors.FgRed(string(*gameengine.GetVersion()))
	fmt.Printf("%s game engine\n\n", gameTitle)
	fmt.Printf("version: %s\n\n", v)
}
