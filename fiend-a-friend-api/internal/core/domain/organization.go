package domain

import (
	"github.com/google/uuid"
)

type Organization struct {
	Id           uuid.UUID
	Address      string
	Name         string
	Email        string
	PasswordHash string
}
