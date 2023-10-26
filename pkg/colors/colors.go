package colors

import "fmt"

type Color struct {
	fgCode int
	bgCode int
}

var (
	red = Color{
		fgCode: 196,
		bgCode: 1,
	}
)

func FgRed(s string) string {
	return fmt.Sprintf("\u001B[38;5;%dm%s\u001B[0m", red.fgCode, s)
}

func BgRed(s string) string {
	return fmt.Sprintf("\u001B[48;5;%dm%s\u001B[0m", red.bgCode, s)
}
