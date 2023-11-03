package common

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID        uuid.UUID
	StartTime time.Time
	LastPOL   time.Time
	Token     uuid.UUID
}

func NewSession() *Session {
	startTime := time.Now()
	return &Session{
		ID:        uuid.New(),
		StartTime: startTime,
		LastPOL:   startTime,
		Token:     uuid.New(),
	}
}
