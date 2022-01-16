package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID
	EmployeeID uuid.UUID
	MenuID     uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewTransaction() *Transaction {
	return &Transaction{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
