package gameclient

import (
	"github.com/h-abranches-dev/connect-4/pkg/versions"
)

var (
	version = new(versions.Version)
)

func GetVersion() string {
	return version.Tag
}

func SetVersion(nv versions.Version) {
	*version = nv
}
