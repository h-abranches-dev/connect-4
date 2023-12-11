package colors

import "fmt"

type Color struct {
	FgCode int
	BgCode int
}

var (
	red = Color{
		FgCode: 196,
		BgCode: 1,
	}

	yellow = Color{
		BgCode: 226,
	}

	blue = Color{
		BgCode: 21,
	}

	black = Color{
		BgCode: 0,
	}
)

func FgRed(s string) string {
	return fmt.Sprintf("\u001B[38;5;%dm%s\u001B[0m", red.FgCode, s)
}

func BgRed(s string) string {
	return fmt.Sprintf("\u001B[48;5;%dm%s\u001B[0m", red.BgCode, s)
}

func Red() Color {
	return red
}

func Yellow() Color {
	return yellow
}

func Blue() Color {
	return blue
}

func Black() Color {
	return black
}
