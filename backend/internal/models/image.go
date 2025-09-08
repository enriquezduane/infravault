package models

import "time"

type Image struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Filename    string    `json:"filename"`
	UploadedAt  time.Time `json:"uploaded_at"`
}
