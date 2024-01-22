package entity

import (
	"time"
)

type User struct {
	ID           string
	PhoneNumber  string
	FullName     string
	PasswordHash string
	CreatedAt    time.Time
}
