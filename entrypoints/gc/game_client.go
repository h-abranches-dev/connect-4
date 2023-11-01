package main

import (
	"fmt"
	"github.com/h-abranches-dev/connect-4/game-client"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
)

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
