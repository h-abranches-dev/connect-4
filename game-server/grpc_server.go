package gameserver

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	game "github.com/h-abranches-dev/connect-4/common"
	gameengine "github.com/h-abranches-dev/connect-4/game-engine"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"log"
	"slices"
	"sync"
	"time"
)

type GameServer struct {
	UnimplementedRouteServer
}

const (
	sessionLastPOLExpirationInterval = 4 * time.Second

	playTimeout = 10000 * time.Millisecond

	maxSessions = 6

	player1 = "P1"
	player2 = "P2"
)

var (
	openSessions      = new([]*GCSession)
	openSessionsMutex sync.Mutex

	activeMatches = new([]*Match)
)

func NewGameServer() *GameServer {
	return &GameServer{}
}

func (gs *GameServer) VerifyCompatibility(ctx context.Context,
	payload *VerifyCompatibilityPayload) (*VerifyCompatibilityResponse, error) {

	gsv := version

	gcv := versions.Version{
		Tag: payload.GameClientVersion,
	}
	if !gsv.Supports(gcv) {
		return nil, utils.FormatErrors([]string{
			fmt.Sprintf("game client version %q is not compatible with the game server version %q.",
				gcv.Tag, gsv.Tag),
			fmt.Sprintf("game client versions compatible with game server version %q: %s", gsv.Tag,
				versions.GetGameClientsVersions(gsv.SupportedVersions)),
		})
	}

	return &VerifyCompatibilityResponse{GameServerVersion: gsv.Tag}, nil
}

func removeGCSession(ss *[]*GCSession, token string) {
	*ss = slices.DeleteFunc(*ss, func(s *GCSession) bool {
		if s.Token.String() == token {
			s.RemoveGCSessionFromMatch()
			return true
		}
		return false
	})
}

func (gs *GameServer) Login(ctx context.Context, in *LoginPayload) (*LoginResponse, error) {
	if openSessions != nil && len(*openSessions) >= maxSessions {
		return nil, fmt.Errorf(game.MaxConnToGSExceededErr)
	}

	ns := NewSession()
	if !slices.Contains(*openSessions, ns) {
		openSessionsMutex.Lock()
		*openSessions = append(*openSessions, ns)

		m := GetMatches(*activeMatches, *ns)

		if m == nil {
			if len(*openSessions) == 1 &&
				len(*activeMatches) == 0 || len(*openSessions)/2 == len(*activeMatches) &&
				len(*openSessions) == len(*activeMatches)*2+1 {

				nm := NewMatch(ns)
				*activeMatches = append(*activeMatches, nm)
			} else {
				m = GetAvailableMatch(*activeMatches)
				if m != nil && m.P2 == nil {
					m.P2 = NewPlayer("P2", ns)
					ns.Match = m
					m.StatusCode = Started
				}
			}
		}

		openSessionsMutex.Unlock()
	}

	log.Printf(">>> number of active game clients: %d\n\n", len(*openSessions))

	return &LoginResponse{
		SessionToken: ns.Token.String(),
	}, nil
}

func updateGCSessionLastPOLTimestamp(token uuid.UUID) {
	sIdx := slices.IndexFunc(*openSessions, func(s *GCSession) bool {
		return s.Token == token
	})
	if sIdx > -1 {
		(*openSessions)[sIdx].LastPOL = time.Now()
		log.Printf(">>> last POL of session %s updated!\n\n", (*openSessions)[sIdx].Token.String())
	}
}

func (gs *GameServer) POL(ctx context.Context, in *POLPayload) (*POLResponse, error) {
	token, err := uuid.Parse(in.SessionToken)
	if err != nil {
		return nil, fmt.Errorf(game.InvalidSessionTokenFormatErr)
	}

	session := GetGCSessions(*openSessions, token)
	if session == nil {
		return nil, fmt.Errorf(game.InvalidSessionTokenErr)
	}

	updateGCSessionLastPOLTimestamp(token)

	if m := GetMatches(*activeMatches, *session); m != nil {
		return &POLResponse{
			MatchStatus: string(m.StatusCode),
		}, nil
	}

	return nil, fmt.Errorf(game.MatchAbandonedErr)
}

func (gs *GameServer) CheckMatchCanStart(ctx context.Context,
	in *CheckMatchCanStartPayload) (*CheckMatchCanStartResponse, error) {

	var match *Match
	sessionToken, err := uuid.Parse(in.SessionToken)
	if err != nil {
		return nil, fmt.Errorf(game.InvalidSessionTokenFormatErr)
	}
	session := GetGCSessions(*openSessions, sessionToken)
	if session == nil {
		return nil, fmt.Errorf(game.InvalidSessionTokenErr)
	}

	canStart := false

	match = GetMatches(*activeMatches, *session)
	if match != nil && match.IsComplete() {
		canStart = true
	}

	return &CheckMatchCanStartResponse{
		CanStart: canStart,
	}, nil
}

func getBoard(boardID uuid.UUID) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := rc.ServeBoard(ctx, &gameengine.ServeBoardPayload{
		BoardID: boardID.String(),
	})
	if err != nil {
		return nil, err
	}
	board := resp.GetBoard()
	return &board, nil
}

func getPlayerID(match Match, session GCSession) (*string, error) {
	plid := new(string)
	if match.P1 != nil && match.P1.Session.Token == session.Token {
		*plid = player1
	} else if match.P2 != nil && match.P2.Session.Token == session.Token {
		*plid = player2
	} else {
		return nil, fmt.Errorf(game.CannotFindPlayerIDErr)
	}
	return plid, nil
}

func (gs *GameServer) ServeBoard(ctx context.Context, in *ServeBoardPayload) (*ServeBoardResponse, error) {
	var board *string
	var playerID *string
	var err error

	sessionToken, err := uuid.Parse(in.SessionToken)
	if err != nil {
		return nil, fmt.Errorf(game.InvalidSessionTokenFormatErr)
	}

	currentSession := GetGCSessions(*openSessions, sessionToken)
	if currentSession == nil {
		return nil, fmt.Errorf(game.InvalidSessionTokenErr)
	}

	match := GetMatches(*activeMatches, *currentSession)
	if match == nil {
		return nil, fmt.Errorf(game.CannotFindMatchErr)
	}

	playerID, err = getPlayerID(*match, *currentSession)
	if err != nil {
		return nil, err
	}

	board, err = getBoard(match.BoardID)
	if err != nil {
		return nil, err
	}

	return &ServeBoardResponse{
		Board:    *board,
		PlayerID: *playerID,
	}, nil
}

func getPlayer(ms []*Match, sessionToken uuid.UUID) *Player {
	mIdx := slices.IndexFunc(ms, func(m *Match) bool {
		return m.P1.Session != nil && m.P1.Session.Token == sessionToken ||
			m.P2.Session != nil && m.P2.Session.Token == sessionToken
	})
	if mIdx == -1 {
		return nil
	}
	matchFound := ms[mIdx]
	if matchFound.P1 != nil && matchFound.P1.Session != nil && matchFound.P1.Session.Token == sessionToken {
		return matchFound.P1
	}
	if matchFound.P2 != nil && matchFound.P2.Session != nil && matchFound.P2.Session.Token == sessionToken {
		return matchFound.P2
	}
	return nil
}

func play(playerCode string, boardID uuid.UUID, column int32) (*bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), playTimeout)
	defer cancel()
	resp, err := rc.Play(ctx, &gameengine.PlayPayload{
		BoardID:    boardID.String(),
		PlayerCode: playerCode,
		Column:     column,
	})
	if err != nil {
		return nil, err
	}

	return &resp.ThereIsWinner, nil
}

func (gs *GameServer) Play(ctx context.Context, in *PlayPayload) (*PlayResponse, error) {
	sessionToken, err := uuid.Parse(in.SessionToken)
	if err != nil {
		return &PlayResponse{
			ThereIsWinner: true,
			Error:         game.InvalidSessionTokenFormatErr,
		}, nil
	}

	currentSession := GetGCSessions(*openSessions, sessionToken)
	if currentSession == nil {
		return &PlayResponse{
			ThereIsWinner: true,
			Error:         game.InvalidSessionTokenErr,
		}, nil
	}

	match := GetMatches(*activeMatches, *currentSession)
	if match == nil {
		return &PlayResponse{
			ThereIsWinner: true,
			Error:         game.CannotFindMatchErr,
		}, nil
	}

	player := getPlayer(*activeMatches, sessionToken)
	if player == nil {
		return &PlayResponse{
			ThereIsWinner: true,
			Error:         game.CannotFindPlayerErr,
		}, nil
	}

	thereIsWinner, err := play(player.Code, match.BoardID, in.Column)
	if err != nil {
		return &PlayResponse{
			ThereIsWinner: true,
			Error:         err.Error(),
		}, nil
	}

	match.SessionLastPlay = currentSession
	if *thereIsWinner {
		if player.Code == "P1" {
			match.StatusCode = WinnerP1
		} else {
			match.StatusCode = WinnerP2
		}
	}

	return &PlayResponse{
		ThereIsWinner: *thereIsWinner,
		Error:         "",
	}, nil
}

func (gs *GameServer) CheckBoardUpdated(ctx context.Context,
	in *CheckBoardUpdatedPayload) (*CheckBoardUpdatedResponse, error) {

	var sessionToken uuid.UUID
	var err error
	sessionToken, err = uuid.Parse(in.SessionToken)
	if err != nil {
		return nil, fmt.Errorf(game.InvalidSessionTokenFormatErr)
	}

	currentSession := GetGCSessions(*openSessions, sessionToken)
	if currentSession == nil {
		return nil, fmt.Errorf(game.InvalidSessionTokenErr)
	}

	match := GetMatches(*activeMatches, *currentSession)
	if match != nil {
		var cp *Player
		if match.P1 != nil && match.P1.Session.Token == currentSession.Token {
			cp = match.P1
		} else if match.P2 != nil && match.P2.Session.Token == currentSession.Token {
			cp = match.P2
		}

		if match.IsFinished() {
			if cp.Code == player1 && match.P2 == nil || cp.Code == player2 && match.P1 == nil {
				return &CheckBoardUpdatedResponse{
					Feedback: "RA",
				}, nil
			}
		}
	}

	var feedback string

	if match != nil && match.SessionLastPlay != nil && match.SessionLastPlay != currentSession {
		feedback = "P"

		if match.StatusCode == WinnerP1 || match.StatusCode == WinnerP2 {
			feedback = "W"
		}
	}

	return &CheckBoardUpdatedResponse{
		Feedback: feedback,
	}, nil
}
