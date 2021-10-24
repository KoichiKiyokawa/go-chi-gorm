package entity

import (
	"time"

	"gorm.io/gorm"
)

// Post example
type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt" example:"2021-10-24T00:12:39.469332+09:00"`
	UpdatedAt time.Time      `json:"updatedAt" example:"2021-10-24T00:12:39.469332+09:00"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Title     string         `json:"title" example:"title"`
	Body      string         `json:"body" example:"body"`
	UserID    uint           `json:"userId"`
	User      User
}
