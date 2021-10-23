package main

import (
	"go-mux-gorm/entity"
	"go-mux-gorm/seed/seeds"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.User{})

	seeds.SeedUsers(db)
}
