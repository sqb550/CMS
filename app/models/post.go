package models

import "time"

type Post struct {
	ID        uint   `json:"id"`
	UserID    int `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"time"`
	
}
