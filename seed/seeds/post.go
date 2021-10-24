package seeds

import (
	"fmt"
	"go-chi-gorm/entity"

	"gorm.io/gorm"
)

func SeedPosts(db *gorm.DB) {
	var posts []entity.Post
	for userIndex := 0; userIndex < 10; userIndex++ {
		for postIndex := 0; postIndex < 100; postIndex++ {

			posts = append(posts, entity.Post{Title: fmt.Sprintf("user%d-title%d", userIndex, postIndex), Body: fmt.Sprintf("user%d-body%d", userIndex, postIndex), UserID: uint(userIndex)})
		}
	}
	db.Create(posts)
}
