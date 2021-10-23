package entity

import (
	"time"

	"gorm.io/gorm"
)

// User example
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt" example:"2021-10-24T00:12:39.469332+09:00"`
	UpdatedAt time.Time      `json:"updatedAt" example:"2021-10-24T00:12:39.469332+09:00"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name" example:"Taro Yamada"`
	Email     string         `json:"email" example:"foo@example.com"`
	Password  string         `json:"-"`
}
