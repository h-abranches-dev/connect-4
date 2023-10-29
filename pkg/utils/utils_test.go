package utils

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestListSrvAddr(t *testing.T) {
	got := ListSrvAddr(33)
	want := "0.0.0.0:33"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestNewAddress(t *testing.T) {
	got := NewAddress("33.33.33.33", 33)
	want := "33.33.33.33:33"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestFormatErrors(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		got := FormatErrors(nil)
		got2 := FormatErrors([]string{})
		want := errors.New("{  }")

		if got.Error() != want.Error() {
			t.Errorf("got %s want %s", got.Error(), want.Error())
		}

		if got2.Error() != want.Error() {
			t.Errorf("got %s want %s", got2, want)
		}
	})

	t.Run("one or more errors", func(t *testing.T) {
		errs := []string{"err1", "err2", "err3"}
		got := FormatErrors(errs)
		want := fmt.Errorf("{\n\t%s\n}", strings.Join(errs, "\n\t"))

		if got.Error() != want.Error() {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestFormatSliceStrings(t *testing.T) {
	t.Run("no slice or empty slice", func(t *testing.T) {
		got := FormatSliceStrings(nil)
		got2 := FormatSliceStrings([]string{})
		want := "[  ]"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		if got2 != want {
			t.Errorf("got %s want %s", got2, want)
		}
	})

	t.Run("slice with one element", func(t *testing.T) {
		got := FormatSliceStrings([]string{"x"})
		want := "[ x ]"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("slice with more than one element", func(t *testing.T) {
		got := FormatSliceStrings([]string{"x", "y", "z"})
		want := "[ x, y, z ]"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestMatchRegex(t *testing.T) {
	t.Run("matches", func(t *testing.T) {
		regex := "^[a]{1}$"
		s := "a"
		got := MatchRegex(regex, s)

		if !got {
			t.Errorf("%q doesn't match the regex %q", s, regex)
		}
	})
}
