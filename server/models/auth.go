package models

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	ID        uuid.UUID
	RoleID    int
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAuth() *Auth {
	return &Auth{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
