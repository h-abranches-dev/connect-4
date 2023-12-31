package versions

import (
	"fmt"
	game "github.com/h-abranches-dev/connect-4/common"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"strings"
)

type Version struct {
	Tag               string
	SupportedVersions []*Version
}

const (
	vvRegex = "^g[s|e|c]-v[0-9]{1,2}.[0-9]{1,3}.[0-9]{1,4}$"
)

var (
	gcV001 = Version{
		Tag: "gc-v0.0.1",
	}
	gsV001 = Version{
		Tag: "gs-v0.0.1",
		SupportedVersions: []*Version{
			&gcV001,
		},
	}
	geV001 = Version{
		Tag: "ge-v0.0.1",
		SupportedVersions: []*Version{
			&gsV001,
		},
	}

	versions = []Version{gcV001, gsV001, geV001}
)

func Get() []Version {
	return versions
}

func GetVersion(vs []Version, vt string) (*Version, error) {
	vp := Version{
		Tag: vt,
	}
	if vp.isValid() {
		for _, v := range vs {
			if v.Tag == vp.Tag {
				if v.SupportedVersions != nil {
					for _, sv := range v.SupportedVersions {
						if !sv.isValid() {
							return nil,
								fmt.Errorf(game.VersionAssociatedNotValidErr,
									v.Tag, sv.Tag)
						}
					}
				}
				return &v, nil
			}
		}
	} else {
		return nil, fmt.Errorf(game.VersionNotValidErr, vt)
	}
	return nil, fmt.Errorf(game.VersionNotFoundErr, vt)
}

func (v Version) isValid() bool {
	valid := utils.MatchRegex(vvRegex, v.Tag)
	if !valid {
		return false
	}
	if v.SupportedVersions != nil {
		for _, vv := range v.SupportedVersions {
			valid = utils.MatchRegex(vvRegex, vv.Tag)
			if !valid {
				return false
			}
		}
	}
	return true
}

// Supports check if version v supports other version ov
func (v Version) Supports(ov Version) bool {
	if v.SupportedVersions != nil {
		for _, sv := range v.SupportedVersions {
			if ov.Tag == sv.Tag {
				return true
			}
		}
	}
	return false
}

func GetSystemVersions(versions []*Version, system string) string {
	var res []string
	for _, v := range versions {
		if strings.Contains(v.Tag, system) {
			res = append(res, v.Tag)
		}
	}
	return utils.FormatSliceStrings(res)
}

func GetGameServersVersions(versions []*Version) string {
	return GetSystemVersions(versions, "gs")
}

func GetGameClientsVersions(versions []*Version) string {
	return GetSystemVersions(versions, "gc")
}
