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
	defaultPort   = 50052
	geDefaultHost = "127.0.0.1"
	geDefaultPort = 50051
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

	displayVersion()
	displayGEAddr()

	conn, rc := gameserver.OpenNewConn(geAddr)
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
	gameserver.RegisterRouteServer(grpcServer, gameserver.NewGameServer())

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

func setGEAddr() {
	geAddr = utils.NewAddress(*geHost, *gePort)
}

func displayGEAddr() {
	fmt.Printf("game engine address: %s\n\n", geAddr)
}
