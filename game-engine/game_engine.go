package gameengine

import (
	"context"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
)

type GameEngine struct {
	UnimplementedRouteServer
}

var (
	version = new(versions.Version)
)

func NewGameEngine() *GameEngine {
	return &GameEngine{}
}

func (ge *GameEngine) Ping(ctx context.Context, payload *PingPayload) (*PingResponse, error) {
	return &PingResponse{Pong: "pong"}, nil
}

func GetVersion() *versions.Version {
	return version
}

func SetVersion(v *versions.Version, nv versions.Version) {
	*v = nv
}
