package versions

import "fmt"

type Version string

func Set(v string) (*Version, error) {
	if v == "" || v == "unknown" {
		return nil, fmt.Errorf("version not set")
	}
	pV := new(Version)
	*pV = Version(v)
	return pV, nil
}
