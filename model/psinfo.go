package model

import (
	"time"
)

type PasswordEntry struct {
	ID        int
	SerialNum string
	Username  string
	Password  string
	CreatedAt time.Time
}

func (PasswordEntry) TableName() string {
	return "passwordEntry"
}
