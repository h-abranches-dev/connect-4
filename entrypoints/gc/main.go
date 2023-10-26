package main

import (
	"fmt"
	"github.com/h-abranches-dev/connect-4/game-client"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"github.com/version-go/ldflags"
	"strings"
)

func main() {
	setVersion(ldflags.Version())

	conn, rc := gameclient.OpenNewConn(utils.NewAddress("127.0.0.1", 50052))
	defer utils.CloseConn(conn)

	start(rc)
}

func setVersion(v string) {
	pv, err := versions.Set(v)
	if err != nil {
		panic(err)
	}
	gameclient.SetVersion(gameclient.GetVersion(), *pv)
}

func displayVersion() {
	v := colors.FgRed(string(*gameclient.GetVersion()))
	fmt.Printf("version: %s\n\n", v)
}

func clearConsole() {
	fmt.Printf("\033[2J")
}

func initialize() {
	gameTitle := colors.FgRed("CONNECT-4")
	fmt.Printf("\n%s\n\n", gameTitle)

	displayVersion()
}

func start(rc gameserver.RouteClient) {
	clearConsole()

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
		clearConsole()

		fmt.Printf("invalid option\n\n")

		fmt.Printf("Write 'START': ")
		_, err = fmt.Scanln(&input)
	}

	fmt.Println()
}
