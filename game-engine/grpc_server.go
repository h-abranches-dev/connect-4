package gameengine

import (
	"context"
	"fmt"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
)

type GameEngine struct {
	UnimplementedRouteServer
}

func NewGameEngine() *GameEngine {
	return &GameEngine{}
}

func (ge *GameEngine) Ping(ctx context.Context, payload *PingPayload) (*PingResponse, error) {
	return &PingResponse{Pong: "pong"}, nil
}

func (ge *GameEngine) VerifyCompatibility(ctx context.Context,
	payload *VerifyCompatibilityPayload) (*VerifyCompatibilityResponse, error) {

	gev := version

	gsv := versions.Version{
		Tag: payload.GameServerVersion,
	}
	if !gev.Supports(gsv) {
		return nil, utils.FormatErrors([]string{
			fmt.Sprintf("game server version %q is not compatible with the game engine version %q.",
				gsv.Tag, gev.Tag),
			fmt.Sprintf("game server versions compatible with game engine version %q: %s", gev.Tag,
				versions.GetGameServersVersions(gev.SupportedVersions)),
		})
	}

	return &VerifyCompatibilityResponse{GameEngineVersion: gev.Tag}, nil
}
