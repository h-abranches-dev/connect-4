package main

import (
	"flag"
	gameengine "github.com/h-abranches-dev/connect-4/game-engine"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/version-go/ldflags"
	"google.golang.org/grpc"
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

	rc gameengine.RouteClient

	errCh = make(chan struct{})
)

func main() {
	flag.Parse()

	gameserver.SetGEAddr(*geHost, *gePort)

	geAddr = gameserver.GetGEAddr()

	if err := gameserver.SetVersion(ldflags.Version()); err != nil {
		utils.PrintError(err)
		return
	}

	var conn *grpc.ClientConn
	var err error
	conn, rc, err = gameserver.OpenNewConnWithGameEngine(geAddr)
	if err != nil {
		utils.PrintError(err)
		return
	}
	defer utils.CloseConn(conn)
	gameserver.SetGERouteClient(rc)

	if err = gameserver.VerifyCompatibility(); err != nil {
		utils.PrintError(err)
		return
	}

	gameserver.DisplayVersion()
	//gameserver.DisplayGEAddr()

	sessionToken, err := gameserver.Connect()
	if err != nil {
		utils.PrintError(err)
		return
	}

	go func() {
		err = gameserver.StartSendingPOL(*sessionToken)
		if err != nil {
			<-errCh
			return
		}
	}()

	go func() {
		if err = gameserver.StartGRPCServer(*port); err != nil {
			utils.PrintError(err)
		}
	}()

	// wait for some error
	errCh <- struct{}{}

	if err != nil {
		utils.PrintError(err)
	}
}
