package versions

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	got := Get()
	want := []Version{
		gcV001, gsV001, gsV002, geV001, gsV003, gsV004, gcV003,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetVersion(t *testing.T) {
	vs := []Version{{
		Tag: "gc-v1.1.1",
		SupportedVersions: []*Version{{
			Tag: "something",
		}},
	}, {
		Tag: "gs-v1.1.1",
		SupportedVersions: []*Version{{
			Tag: "gc-v1.1.2",
		}},
	}, {
		Tag: "gc-v3.3.3",
	}}
	t.Run("version provided not found", func(t *testing.T) {
		vp := "ge-v1.1.1"
		_, err := GetVersion(vs, vp)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})

	t.Run("version provided not valid", func(t *testing.T) {
		vp := "something"
		_, err := GetVersion(vs, vp)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})

	t.Run("version provided has version associated that isn't valid", func(t *testing.T) {
		vp := "gc-v1.1.1"
		_, err := GetVersion(vs, vp)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})

	t.Run("version provided has version associated that is valid", func(t *testing.T) {
		vp := "gs-v1.1.1"
		_, err := GetVersion(vs, vp)

		if err != nil {
			t.Error("didn't get an error but wanted one")
		}
	})

	t.Run("version provided has version associated that is valid", func(t *testing.T) {
		vp := "gc-v3.3.3"
		_, err := GetVersion(vs, vp)
		fmt.Println("err:", err)

		if err != nil {
			t.Error("didn't get an error but wanted one")
		}
	})
}

func TestIsValid(t *testing.T) {
	t.Run("version doesn't match the regex", func(t *testing.T) {
		v := Version{
			Tag: "something",
		}
		got := v.isValid()

		if got {
			t.Errorf("want %t got %t", false, true)
		}
	})

	t.Run("version match the regex and hasn't any supported version associated", func(t *testing.T) {
		v := Version{
			Tag: "gc-v1.1.1",
		}
		got := v.isValid()

		if !got {
			t.Errorf("want %t got %t", true, false)
		}
	})

	t.Run("version matches the regex but at least one supported versions doesn't", func(t *testing.T) {
		v := Version{
			Tag: "gc-v1.1.1",
			SupportedVersions: []*Version{{
				Tag: "something",
			}},
		}
		got := v.isValid()

		if got {
			t.Errorf("want %t got %t", false, true)
		}
	})

	t.Run("version and the supported versions match the regex", func(t *testing.T) {
		v := Version{
			Tag: "gc-v1.1.1",
			SupportedVersions: []*Version{{
				Tag: "gs-v1.1.1",
			}, {
				Tag: "ge-v21.5.33",
			}},
		}
		got := v.isValid()

		if !got {
			t.Errorf("want %t got %t", true, false)
		}
	})
}

func TestSupports(t *testing.T) {
	t.Run("no supported versions", func(t *testing.T) {
		v := Version{
			Tag:               "something",
			SupportedVersions: nil,
		}
		v2 := Version{
			Tag:               "something",
			SupportedVersions: []*Version{},
		}
		ov := Version{
			Tag: "other",
		}
		got := v.Supports(ov)
		got2 := v2.Supports(ov)

		if got {
			t.Errorf("want %t got %t", false, true)
		}

		if got2 {
			t.Errorf("want %t got %t", false, true)
		}
	})

	t.Run("version not supported", func(t *testing.T) {
		v := Version{
			Tag: "something",
			SupportedVersions: []*Version{{
				Tag: "a",
			}},
		}
		ov := Version{
			Tag: "x",
		}
		got := v.Supports(ov)

		if got {
			t.Errorf("want %t got %t", false, true)
		}
	})

	t.Run("version supported", func(t *testing.T) {
		v := Version{
			Tag: "something",
			SupportedVersions: []*Version{{
				Tag: "z",
			}},
		}
		ov := Version{
			Tag: "z",
		}
		got := v.Supports(ov)

		if !got {
			t.Errorf("want %t got %t", true, false)
		}
	})
}

func TestGetSystemVersions(t *testing.T) {
	t.Run("no versions found", func(t *testing.T) {
		want := "[  ]"

		system := "xy"

		vs := []*Version{{
			Tag: "something",
		}}
		got := GetSystemVersions(vs, system)

		var vs2 []*Version
		got2 := GetSystemVersions(vs2, system)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		if got2 != want {
			t.Errorf("got %s want %s", got2, want)
		}
	})

	t.Run("found one version", func(t *testing.T) {
		vs := []*Version{{
			Tag: "xy-v1.1.1",
		}, {
			Tag: "something",
		}}
		system := "xy"
		got := GetSystemVersions(vs, system)
		want := "[ xy-v1.1.1 ]"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("found more than one version", func(t *testing.T) {
		vs := []*Version{{
			Tag: "xy-v1.1.1",
		}, {
			Tag: "xy-v1.1.2",
		}, {
			Tag: "something",
		}, {
			Tag: "foobar",
		}, {
			Tag: "something",
		}, {
			Tag: "xy-v0.0.0",
		}, {
			Tag: "ab-v0.0.0",
		}}
		system := "xy"
		got := GetSystemVersions(vs, system)
		want := "[ xy-v1.1.1, xy-v1.1.2, xy-v0.0.0 ]"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestGetGameServersVersions(t *testing.T) {
	t.Run("game servers versions found", func(t *testing.T) {
		vs := []*Version{{
			Tag: "gs-v1.1.1",
		}, {
			Tag: "something",
		}}
		got := GetGameServersVersions(vs)
		want := "[ gs-v1.1.1 ]"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("game servers versions not found", func(t *testing.T) {
		vs := []*Version{{
			Tag: "something",
		}, {
			Tag: "something2",
		}}
		got := GetGameServersVersions(vs)
		want := "[  ]"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestGetGameClientsVersions(t *testing.T) {
	t.Run("game clients versions found", func(t *testing.T) {
		vs := []*Version{{
			Tag: "gc-v1.1.1",
		}, {
			Tag: "something",
		}}
		got := GetGameClientsVersions(vs)
		want := "[ gc-v1.1.1 ]"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("game clients versions not found", func(t *testing.T) {
		vs := []*Version{{
			Tag: "something",
		}, {
			Tag: "something2",
		}}
		got := GetGameServersVersions(vs)
		want := "[  ]"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
