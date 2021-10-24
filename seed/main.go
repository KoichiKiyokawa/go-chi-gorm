package main

import (
	"go-chi-gorm/entity"
	"go-chi-gorm/seed/seeds"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.User{}, &entity.Post{})

	seeds.SeedUsers(db)
	seeds.SeedPosts(db)
}
