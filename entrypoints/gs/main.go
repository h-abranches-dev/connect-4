package main

import (
	"flag"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/version-go/ldflags"
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
	geAddr = new(string)

	stopGS = make(chan bool)
)

func main() {
	flag.Parse()

	gameserver.SetGEAddr(geAddr, *geHost, *gePort)

	if err := gameserver.SetVersion(ldflags.Version()); err != nil {
		utils.PrintError(err)
		return
	}

	conn, rc, err := gameserver.OpenNewConn(*geAddr)
	if err != nil {
		utils.PrintError(err)
		return
	}
	defer utils.CloseConn(conn)

	if err = gameserver.VerifyCompatibility(rc); err != nil {
		utils.PrintError(err)
		return
	}

	gameserver.DisplayGEAddr(*geAddr)

	sessionToken, err := gameserver.Connect(rc)
	if err != nil {
		utils.PrintError(err)
		return
	}

	go gameserver.SendPOL(rc, *sessionToken, stopGS)

	go func() {
		if err = gameserver.StartGRPCServer(*port); err != nil {
			utils.PrintError(err)
		}
	}()

	<-stopGS
}
