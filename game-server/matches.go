package gameserver

import (
	"fmt"
	"github.com/google/uuid"
	"slices"
)

type MatchStatusCode string

const (
	NotStarted MatchStatusCode = "NOT_STARTED"
	Started    MatchStatusCode = "STARTED"
	WinnerP1   MatchStatusCode = "WINNER_P1"
	WinnerP2   MatchStatusCode = "WINNER_P2"
	Abandoned  MatchStatusCode = "ABANDONED"
)

type Player struct {
	Code    string
	Session *GCSession
}

type Match struct {
	Id              string
	BoardID         uuid.UUID
	P1              *Player
	P2              *Player
	SessionLastPlay *GCSession
	StatusCode      MatchStatusCode
}

func NewPlayer(code string, session *GCSession) *Player {
	return &Player{
		Code:    code,
		Session: session,
	}
}

func NewMatch(session *GCSession) *Match {
	nm := &Match{
		Id:         fmt.Sprintf("M_%s", uuid.NewString()[:8]),
		BoardID:    uuid.New(),
		P1:         NewPlayer("P1", session),
		P2:         nil,
		StatusCode: NotStarted,
	}
	session.Match = nm
	return nm
}

func (m *Match) IsFinished() bool {
	if m.StatusCode == WinnerP1 ||
		m.StatusCode == WinnerP2 ||
		m.StatusCode == Abandoned {
		return true
	}
	return false
}

func GetMatches(matches []*Match, session GCSession) *Match {
	mIdx := slices.IndexFunc(matches, func(m *Match) bool {
		return m.P1 != nil && m.P1.Session != nil && *m.P1.Session == session ||
			m.P2 != nil && m.P2.Session != nil && *m.P2.Session == session
	})
	if mIdx == -1 {
		return nil
	}
	return matches[mIdx]
}

func GetAvailableMatch(matches []*Match) *Match {
	mIdx := slices.IndexFunc(matches, func(m *Match) bool {
		return m.P1 != nil && m.P2 == nil
	})
	if mIdx == -1 {
		return nil
	}
	return matches[mIdx]
}

func (m *Match) IsComplete() bool {
	return m.P1 != nil && m.P2 != nil
}

func RemoveMatch(matches *[]*Match, matchID string) {
	*matches = slices.DeleteFunc(*matches, func(m *Match) bool {
		return m.Id == matchID
	})
}
