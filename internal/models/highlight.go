package models

import (
	"time"

	"github.com/google/uuid"
)

type Highlight struct {
	ID        uuid.UUID `json:"id"`
	ImageURL  string    `json:"image_url"`
	Caption   *string   `json:"caption"`
	Term      *Term     `json:"term"`
	Year      *int      `json:"year"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateHighlightRequest struct {
	ImageURL string  `json:"image_url"`
	Caption  *string `json:"caption"`
	Term     *Term   `json:"term"`
	Year     *int    `json:"year"`
}

func (r CreateHighlightRequest) Validate() string {
	if r.ImageURL == "" {
		return "image_url is required"
	}
	return ""
}
