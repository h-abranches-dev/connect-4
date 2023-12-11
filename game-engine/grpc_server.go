package gameengine

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	game "github.com/h-abranches-dev/connect-4/common"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"log"
	"slices"
	"strings"
	"sync"
	"time"
)

type GameEngine struct {
	UnimplementedRouteServer
}

const (
	checkExpiredGSSessionInterval      = 2 * time.Second
	gsSessionLastPOLExpirationInterval = 4 * time.Second
)

var (
	gsSession *GSSession

	gsSessionMutex sync.Mutex

	matchBoards *[]*Board
)

func NewGameEngine() *GameEngine {
	return &GameEngine{}
}

func (ge *GameEngine) VerifyCompatibility(ctx context.Context,
	payload *VerifyCompatibilityPayload) (*VerifyCompatibilityResponse, error) {

	gev := version

	gsv := versions.Version{
		Tag: payload.GameServerVersion,
	}
	if !gev.Supports(gsv) {
		return nil, utils.FormatErrors([]string{
			fmt.Sprintf(game.GEAndGSVerNotCompatibleErr, gsv.Tag, gev.Tag),
			fmt.Sprintf(game.GEAndGSVerNotCompatibleFeedback, gev.Tag,
				versions.GetGameServersVersions(gev.SupportedVersions)),
		})
	}

	return &VerifyCompatibilityResponse{}, nil
}

func newGSSession() *GSSession {
	ns := NewGSSession()
	gsSessionMutex.Lock()
	gsSession = ns
	gsSessionMutex.Unlock()

	return gsSession
}

func startCheckingIfGSSessionExpired() {
	ticker := time.NewTicker(checkExpiredGSSessionInterval)
	for {
		select {
		case <-ticker.C:
			if gsSession == nil {
				log.Printf(">>> no game server connected\n\n")
				ticker.Stop()
				return
			} else {
				expirationTime := gsSession.LastPOL.Add(gsSessionLastPOLExpirationInterval)
				now := time.Now()
				if expirationTime.Before(now) {
					gsSession = nil
					matchBoards = nil
				}
			}
		}
	}
}

func (ge *GameEngine) Connect(ctx context.Context, in *ConnectPayload) (*ConnectResponse, error) {
	if gsSession != nil {
		return nil, fmt.Errorf(game.MaxConnToGEExceededErr)
	}

	gsSession = newGSSession()

	go startCheckingIfGSSessionExpired()

	log.Printf(">>> game server connected\n\n")

	return &ConnectResponse{
		SessionToken: gsSession.Token.String(),
	}, nil
}

func tokenProvidedIsValid(tokenProvided uuid.UUID) bool {
	if gsSession == nil {
		return false
	}
	return gsSession != nil && gsSession.Token == tokenProvided
}

func updateGSSessionLastPOLTimestamp() {
	gsSession.LastPOL = time.Now()
}

func removeBoard(boardID string) {
	*matchBoards = slices.DeleteFunc(*matchBoards, func(b *Board) bool {
		return b.Id.String() == boardID
	})
}

func removeOutdatedGameBoards(encodedBoards string) {
	if matchBoards != nil {
		encodedBoardsIDs := strings.Split(encodedBoards, ";")
		for _, b := range *matchBoards {
			found := false
			for _, ebid := range encodedBoardsIDs {
				if ebid == b.Id.String() {
					found = true
					break
				}
			}
			if !found {
				removeBoard(b.Id.String())
			}
		}
	}
}

func (ge *GameEngine) POL(ctx context.Context, in *POLPayload) (*POLResponse, error) {
	removeOutdatedGameBoards(in.EncodedBoards)

	token, err := uuid.Parse(in.SessionToken)
	if err != nil {
		return nil, fmt.Errorf(game.InvalidSessionTokenFormatErr)
	}

	if ok := tokenProvidedIsValid(token); !ok {
		return nil, fmt.Errorf(game.InvalidSessionTokenErr)
	}

	updateGSSessionLastPOLTimestamp()

	return &POLResponse{}, nil
}

func getBoard(bs []*Board, boardID uuid.UUID) *Board {
	bIdx := slices.IndexFunc(bs, func(b *Board) bool {
		return b.Id == boardID
	})
	if bIdx == -1 {
		return nil
	}
	return bs[bIdx]
}

func getPalette(palettes []Palette, paletteID string) (*Palette, error) {
	idx := slices.IndexFunc(palettes, func(p Palette) bool {
		return strings.ToUpper(p.Id) == strings.ToUpper(paletteID)
	})
	if idx == -1 {
		return nil, fmt.Errorf(game.CannotFindPaletteErr)
	}
	return &palettes[idx], nil
}

func (ge *GameEngine) ServeBoard(ctx context.Context, in *ServeBoardPayload) (*ServeBoardResponse, error) {
	boardID, err := uuid.Parse(in.BoardID)
	if err != nil {
		return nil, fmt.Errorf(game.InvalidBoardIDFormatErr)
	}

	p, err := getPalette(Palettes, PaletteYRB)
	if err != nil {
		return nil, err
	}

	if matchBoards == nil {
		matchBoards = new([]*Board)
	}
	b := getBoard(*matchBoards, boardID)
	if b == nil {
		b = New(matchBoards, boardID, *p)
	}

	err = b.String()
	if err != nil {
		return nil, fmt.Errorf(game.CannotRenderBoardErr)
	}

	return &ServeBoardResponse{Board: (*b).Output}, nil
}

func (ge *GameEngine) Play(ctx context.Context, in *PlayPayload) (*PlayResponse, error) {
	boardID, err := uuid.Parse(in.BoardID)
	if err != nil {
		return nil, fmt.Errorf(game.InvalidBoardIDFormatErr)
	}

	b := getBoard(*matchBoards, boardID)
	if b == nil {
		return nil, fmt.Errorf(game.InvalidBoardIDErr)
	}

	player, err := GetPlayer(in.PlayerCode)
	if err != nil {
		return nil, err
	}

	thereIsWinner, err := b.Play(*player, in.Column)
	if err != nil {
		return nil, err
	}
	if err = b.String(); err != nil {
		return nil, fmt.Errorf(game.CannotRenderBoardErr)
	}
	return &PlayResponse{
		ThereIsWinner: thereIsWinner,
	}, nil
}
