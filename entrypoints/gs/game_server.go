package main

import (
	"fmt"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
)

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
