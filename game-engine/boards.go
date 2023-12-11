package gameengine

import (
	"fmt"
	"github.com/google/uuid"
	game "github.com/h-abranches-dev/connect-4/common"
	"github.com/h-abranches-dev/connect-4/pkg/colors"
	"github.com/h-abranches-dev/connect-4/pkg/stacks"
	"slices"
)

type Token rune

type Player struct {
	code       string
	playsToken Token
}

type Palette struct {
	Id               string
	gridColor        colors.Color
	emptySpacesColor colors.Color
	p1TokenColor     colors.Color
	p2TokenColor     colors.Color
}

type Board struct {
	Id      uuid.UUID
	Stacks  [numColumns]*stacks.Stack
	Grid    [numRows][numColumns]string
	Output  string
	Palette Palette
	Tokens  []Token
}

const (
	p1Code                = "P1"
	p2Code                = "P2"
	P1PlaysToken    Token = 'A'
	P2PlaysToken    Token = 'B'
	emptySpaceToken Token = 'E'
	numRows               = 6
	numColumns            = 7

	PaletteYRB = "YRB"
)

var (
	Palettes = []Palette{{
		Id:               PaletteYRB,
		gridColor:        colors.Blue(),
		emptySpacesColor: colors.Black(),
		p1TokenColor:     colors.Yellow(),
		p2TokenColor:     colors.Red(),
	}}

	p1 = Player{
		code:       p1Code,
		playsToken: P1PlaysToken,
	}
	p2 = Player{
		code:       p2Code,
		playsToken: P2PlaysToken,
	}
)

func GetPlayer(playerCode string) (*Player, error) {
	switch playerCode {
	case p1.code:
		return &p1, nil
	case p2.code:
		return &p2, nil
	default:
		return nil, fmt.Errorf(game.InvalidPlayerCodeErr)
	}
}

func New(matchBoards *[]*Board, boardID uuid.UUID, palette Palette) *Board {
	board := &Board{
		Id:      boardID,
		Stacks:  [numColumns]*stacks.Stack{},
		Palette: palette,
		Tokens:  []Token{P1PlaysToken, P2PlaysToken, emptySpaceToken},
	}
	for i := 0; i < numColumns; i++ {
		board.Stacks[i] = stacks.New()
	}
	board.updateBoard()

	*matchBoards = append(*matchBoards, board)

	return board
}

func getColumn(st *stacks.Stack) [1][numRows]string {
	res := [1][numRows]string{}
	for i := numRows - 1; i >= 0; i-- {
		if i > len(*st)-1 {
			res[0][i] = fmt.Sprintf("%s", string(emptySpaceToken))
			continue
		}
		res[0][i] = fmt.Sprintf("%s", string((*st)[i]))
	}
	return res
}

func (b *Board) updateBoard() {
	for i := 0; i < numRows; i++ {
		for j := 0; j < numColumns; j++ {
			column := getColumn(b.Stacks[j])[0]
			b.Grid[i][j] = column[i]
		}
	}
}

func (b *Board) addColsIDs() {
	for i := 1; i <= numColumns; i++ {
		b.Output += fmt.Sprintf("\u001B[48;5;%dm   %d  \u001B[0m",
			b.Palette.gridColor.BgCode, i)
	}
	b.Output += "\n\n"
}

func (b *Board) addGridRow() {
	b.Output += fmt.Sprintf("\u001B[48;5;%dm%s\u001B[0m\n",
		b.Palette.gridColor.BgCode, "                                          ")
}

func (b *Board) addEmptySpaces() {
	b.Output += fmt.Sprintf("\u001B[48;5;%dm%s\u001B[0m\u001B[48;5;%dm%s\u001B[0m\u001B[48;5;%dm%s\u001B[0m",
		b.Palette.gridColor.BgCode, "  ", b.Palette.emptySpacesColor.BgCode, "  ", b.Palette.gridColor.BgCode, "  ")
}

func isValidPlaysToken(tokens []Token, t Token) bool {
	if idx := slices.IndexFunc(tokens, func(it Token) bool {
		return t == it
	}); idx == -1 {
		return false
	}
	return t == P1PlaysToken || t == P2PlaysToken
}

func playerCodeColor(p Palette, playerToken Token) (*colors.Color, error) {
	switch playerToken {
	case P1PlaysToken:
		return &p.p1TokenColor, nil
	case P2PlaysToken:
		return &p.p2TokenColor, nil
	default:
		return nil, fmt.Errorf("player token is wrong")
	}
}

func (b *Board) addTokens(playerToken Token) error {
	playerColor, err := playerCodeColor(b.Palette, playerToken)
	if err != nil {
		return err
	}

	b.Output += fmt.Sprintf("\u001B[48;5;%dm%s\u001B[0m\u001B[48;5;%dm%s\u001B[0m\u001B[48;5;%dm%s\u001B[0m",
		b.Palette.gridColor.BgCode, "  ", playerColor.BgCode, "  ", b.Palette.gridColor.BgCode, "  ")

	return nil
}

func (b *Board) String() error {
	b.Output = "\n"
	b.addColsIDs()

	b.addGridRow()
	for i := len(b.Grid) - 1; i >= 0; i-- {
		for j := 0; j < len(b.Grid[i]); j++ {
			t := Token([]rune(b.Grid[i][j])[0])
			if t == emptySpaceToken {
				b.addEmptySpaces()
				continue
			}
			if !isValidPlaysToken(b.Tokens, t) {
				return fmt.Errorf("%q is an invalid player token", string(t))
			}
			err := b.addTokens(t)
			if err != nil {
				return err
			}
		}
		b.Output += "\n"
		b.addGridRow()
	}
	return nil
}

func (b *Board) thereIsWinner() *Token {
	found := false
	var t string
	var pT *Token
	for i := 0; i < numRows; i++ {
		for j := 0; j < numColumns; j++ {
			if b.Grid[i][j] == "E" {
				continue
			}
			t = b.Grid[i][j]
			// horizontal --
			if j+1 < numColumns && b.Grid[i][j+1] == t {
				if j+2 < numColumns && b.Grid[i][j+2] == t {
					if j+3 < numColumns && b.Grid[i][j+3] == t {
						found = true
						break
					}
				}
			}
			// vertical |
			if i+1 < numRows && b.Grid[i+1][j] == t {
				if i+2 < numRows && b.Grid[i+2][j] == t {
					if i+3 < numRows && b.Grid[i+3][j] == t {
						found = true
						break
					}
				}
			}
			// diagonal /
			if i+1 < numRows && j+1 < numColumns && b.Grid[i+1][j+1] == t {
				if i+2 < numRows && j+2 < numColumns && b.Grid[i+2][j+2] == t {
					if i+3 < numRows && j+3 < numColumns && b.Grid[i+3][j+3] == t {
						found = true
						break
					}
				}
			}

			// diagonal \
			if i-1 >= 0 && j+1 < numColumns && b.Grid[i-1][j+1] == t {
				if i-2 >= 0 && j+2 < numColumns && b.Grid[i-2][j+2] == t {
					if i-3 >= 0 && j+3 < numColumns && b.Grid[i-3][j+3] == t {
						found = true
						break
					}
				}
			}
		}
		if found {
			pT = new(Token)
			*pT = Token([]rune(t)[0])
			break
		}
	}
	return pT
}

func (b *Board) Play(p Player, c int32) (bool, error) {
	if c < 1 || c > numColumns {
		return false, fmt.Errorf(game.InvalidColumnNumberPlayErr)
	}
	if len(*b.Stacks[c-1]) >= numRows {
		return false, fmt.Errorf(game.ColumnIsFullPlayErr)
	}
	b.Stacks[c-1].Push(rune(p.playsToken))
	b.updateBoard()
	t := b.thereIsWinner()
	if t != nil {
		return true, nil
	}
	return false, nil
}
