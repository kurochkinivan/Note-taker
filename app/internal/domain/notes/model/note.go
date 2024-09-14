package model

import "time"

type Note struct {
	ID        string    `json:"-"`
	UserID    string    `json:"-"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
