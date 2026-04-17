package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Permission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}
