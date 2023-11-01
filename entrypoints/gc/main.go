package main

import (
	"flag"
	"fmt"
	"github.com/h-abranches-dev/connect-4/game-client"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/version-go/ldflags"
	"strings"
)

const (
	gsDefaultHost = "127.0.0.1"
	gsDefaultPort = 50052
)

var (
	gsHost = flag.String("gsHost", gsDefaultHost, "the game server host")
	gsPort = flag.Int("gsPort", gsDefaultPort, "the game server port")
	gsAddr string
)

func main() {
	flag.Parse()

	setGSAddr()
	setVersion(ldflags.Version())

	conn, rc := gameclient.OpenNewConn(gsAddr)
	defer utils.CloseConn(conn)

	verifyCompatibility(rc)

	start(rc)
}

func initialize() {
	gameTitle := colors.FgRed("CONNECT-4")
	fmt.Printf("\n%s\n\n", gameTitle)

	displayVersion()
	displayGSAddr()
}

func start(rc gameserver.RouteClient) {
	utils.ClearConsole()

	initialize()

	err := gameclient.PingGameServer(rc)
	if err != nil {
		fmt.Printf("%s\n\n", err.Error())
		return
	}

	fmt.Printf("Write 'START': ")
	var input string
	_, err = fmt.Scanln(&input)
	for err != nil || strings.ToUpper(input) != "START" {
		utils.ClearConsole()

		fmt.Printf("invalid option\n\n")

		fmt.Printf("Write 'START': ")
		_, err = fmt.Scanln(&input)
	}

	fmt.Println()
}
