package gameclient

import "github.com/h-abranches-dev/connect-4/pkg/versions"

var (
	gameClientVersion = new(versions.Version)
)

func GetVersion() *versions.Version {
	return gameClientVersion
}

func SetVersion(v *versions.Version, nv versions.Version) {
	*v = nv
}
