package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `json:"id"`
	UserID    int            `json:"user_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"time"`
	UpdatedAt time.Time      `json:"updated_time"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Likes     int            `json:"likes"`
}
