package main

import (
	"flag"
	"fmt"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"github.com/version-go/ldflags"
	"google.golang.org/grpc"
	"net"
)

const (
	defaultPort = 50052
)

var (
	port = flag.Int("port", defaultPort, "the game server port")
)

func main() {
	setVersion(ldflags.Version())

	displayVersion()

	flag.Parse()

	conn, rc := gameserver.OpenNewConn(utils.NewAddress("127.0.0.1", 50051))
	defer utils.CloseConn(conn)

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
	gameserver.RegisterRouteServer(grpcServer, gameserver.NewGameServerWrapper())

	fmt.Printf("game server is listening on the address %s\n", colors.FgRed(lis.Addr().String()))

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
	gameserver.SetVersion(gameserver.GetVersion(), *pv)
}

func displayVersion() {
	gameTitle := colors.FgRed("CONNECT-4")
	v := colors.FgRed(string(*gameserver.GetVersion()))
	fmt.Printf("%s game server\n\n", gameTitle)
	fmt.Printf("version: %s\n\n", v)
}
