package main

import (
	"fmt"
	"github.com/h-abranches-dev/connect-4/gameclient"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"github.com/version-go/ldflags"
	"strings"
)

func main() {
	setVersion(ldflags.Version())

	start()
}

func setVersion(v string) {
	pv, err := versions.Set(v)
	if err != nil {
		panic(err)
	}
	gameclient.SetVersion(gameclient.GetVersion(), *pv)
}

func displayVersion() {
	v := string(*gameclient.GetVersion())
	fmt.Printf("version: %s\n\n", v)
}

func clearConsole() {
	fmt.Printf("\033[2J")
}

func initialize() {
	gameTitle := "CONNECT-4"
	fmt.Printf("\n%s\n\n", gameTitle)

	displayVersion()
}

func start() {
	clearConsole()

	initialize()

	fmt.Printf("Write 'START': ")
	var input string
	_, err := fmt.Scanln(&input)
	for err != nil || strings.ToUpper(input) != "START" {
		clearConsole()

		fmt.Printf("invalid option\n")

		initialize()

		fmt.Printf("Write 'START': ")
		_, err = fmt.Scanln(&input)
	}
}
