package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/h-abranches-dev/connect-4/game-client"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"github.com/version-go/ldflags"
	"strings"
	"time"
)

const (
	gsDefaultHost = "127.0.0.1"
	gsDefaultPort = 50052

	verifyCompatibilityTimeout = 500 * time.Millisecond
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

func clearConsole() {
	fmt.Printf("\033[2J")
}

func setVersion(v string) {
	pv, err := versions.GetVersion(versions.Get(), v)
	if err != nil {
		panic(err)
	}
	gameclient.SetVersion(*pv)
}

func displayVersion() {
	v := colors.FgRed(gameclient.GetVersion())
	fmt.Printf("version: %s\n\n", v)
}

func setGSAddr() {
	gsAddr = utils.NewAddress(*gsHost, *gsPort)
}

func displayGSAddr() {
	fmt.Printf("game server address: %s\n\n", gsAddr)
}

func verifyCompatibility(rc gameserver.RouteClient) {
	gcv := gameclient.GetVersion()

	ctx, cancel := context.WithTimeout(context.Background(), verifyCompatibilityTimeout)
	defer cancel()

	_, err := rc.VerifyCompatibility(ctx, &gameserver.VerifyCompatibilityPayload{
		GameClientVersion: gcv,
	})
	if err != nil {
		panic(err)
	}

	displayVersion()
}
