package models

import "time"

type Role struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Permission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}
