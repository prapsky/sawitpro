package entity

import (
	"time"
)

type User struct {
	ID               uint64
	PhoneNumber      string
	FullName         string
	PasswordHash     string
	SuccessfulLogins uint64
	CreatedAt        time.Time
	LastLoginAt      time.Time
}
