package main

import (
	"context"
	gameengine "github.com/h-abranches-dev/connect-4/game-engine"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"time"
)

const (
	verifyCompatibilityTimeout = 500 * time.Millisecond
)

func verifyCompatibility(rc gameengine.RouteClient) {
	gsv := gameserver.GetVersion()

	ctx, cancel := context.WithTimeout(context.Background(), verifyCompatibilityTimeout)
	defer cancel()

	_, err := rc.VerifyCompatibility(ctx, &gameengine.VerifyCompatibilityPayload{
		GameServerVersion: gsv,
	})
	if err != nil {
		panic(err)
	}

	displayVersion()
}
