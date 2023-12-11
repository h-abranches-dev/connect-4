package gameengine

import (
	"github.com/google/uuid"
	"time"
)

type GSSession struct {
	ID        uuid.UUID
	StartTime time.Time
	LastPOL   time.Time
	Token     uuid.UUID
}

func NewGSSession() *GSSession {
	startTime := time.Now()
	return &GSSession{
		ID:        uuid.New(),
		StartTime: startTime,
		LastPOL:   startTime,
		Token:     uuid.New(),
	}
}
