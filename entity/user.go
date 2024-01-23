package entity

import (
	"time"
)

type User struct {
	ID           uint64
	PhoneNumber  string
	FullName     string
	PasswordHash string
	CreatedAt    time.Time
}
