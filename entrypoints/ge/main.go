package main

import (
	"flag"
	gameengine "github.com/h-abranches-dev/connect-4/game-engine"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/version-go/ldflags"
)

const (
	defaultPort = 50051
)

var (
	port = flag.Int("port", defaultPort, "the game engine port")
)

func main() {
	flag.Parse()

	if err := gameengine.SetVersion(ldflags.Version()); err != nil {
		utils.PrintError(err)
		return
	}

	gameengine.DisplayVersion()

	if err := gameengine.StartGRPCServer(*port); err != nil {
		utils.PrintError(err)
		return
	}
}
