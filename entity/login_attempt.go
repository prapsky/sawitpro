package entity

import (
	"time"
)

type LoginAttempt struct {
	ID          uint64
	UserID      uint64
	Success     bool
	AttemptedAt time.Time
}
