package entity

import (
	"time"

	"gorm.io/gorm"
)

// User example
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
}
