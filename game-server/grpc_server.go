package gameserver

import (
	"context"
	"fmt"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
)

type GameServer struct {
	UnimplementedRouteServer
}

func NewGameServer() *GameServer {
	return &GameServer{}
}

func (gs *GameServer) Ping(ctx context.Context, payload *PingPayload) (*PingResponse, error) {
	return &PingResponse{Pong: "pong"}, nil
}

func (gs *GameServer) VerifyCompatibility(ctx context.Context,
	payload *VerifyCompatibilityPayload) (*VerifyCompatibilityResponse, error) {

	gsv := version

	gcv := versions.Version{
		Tag: payload.GameClientVersion,
	}
	if !gsv.Supports(gcv) {
		return nil, utils.FormatErrors([]string{
			fmt.Sprintf("game client version %q is not compatible with the game server version %q.",
				gcv.Tag, gsv.Tag),
			fmt.Sprintf("game client versions compatible with game server version %q: %s", gsv.Tag, versions.GetGameClientsVersions(gsv.SupportedVersions)),
		})
	}

	return &VerifyCompatibilityResponse{GameServerVersion: gsv.Tag}, nil
}
