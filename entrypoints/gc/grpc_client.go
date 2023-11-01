package main

import (
	"context"
	"github.com/h-abranches-dev/connect-4/game-client"
	gameserver "github.com/h-abranches-dev/connect-4/game-server"
	"time"
)

const (
	verifyCompatibilityTimeout = 500 * time.Millisecond
)

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
