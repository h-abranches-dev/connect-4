package gameserver

import (
	"github.com/google/uuid"
	"slices"
	"time"
)

type GCSession struct {
	ID        uuid.UUID
	StartTime time.Time
	LastPOL   time.Time
	Token     uuid.UUID
	Match     *Match
}

func NewSession() *GCSession {
	startTime := time.Now()
	return &GCSession{
		ID:        uuid.New(),
		StartTime: startTime,
		LastPOL:   startTime,
		Token:     uuid.New(),
	}
}

func GetGCSessions(sessions []*GCSession, sessionToken uuid.UUID) *GCSession {
	sIdx := slices.IndexFunc(sessions, func(s *GCSession) bool {
		return s.Token == sessionToken
	})
	if sIdx == -1 {
		return nil
	}
	return sessions[sIdx]
}

func (s *GCSession) RemoveGCSessionFromMatch() {
	m := s.Match
	if m != nil {
		if m.P1 != nil && m.P1.Session.Token == s.Token {
			m.P1 = nil
		}
		if m.P2 != nil && m.P2.Session.Token == s.Token {
			m.P2 = nil
		}
	}
}
