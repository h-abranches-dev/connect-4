package main

import (
	"fmt"
	gameengine "github.com/h-abranches-dev/connect-4/game-engine"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
)

func setVersion(v string) {
	pv, err := versions.GetVersion(versions.Get(), v)
	if err != nil {
		panic(err)
	}
	gameengine.SetVersion(*pv)
}

func displayVersion() {
	gameTitle := colors.FgRed("CONNECT-4")
	v := colors.FgRed(gameengine.GetVersion())
	fmt.Printf("%s game engine\n\n", gameTitle)
	fmt.Printf("version: %s\n\n", v)
}
