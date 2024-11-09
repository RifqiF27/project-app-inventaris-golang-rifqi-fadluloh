package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

type Session struct {
	SessionID uuid.UUID
	UserID    int
	LoginTime time.Time
}
