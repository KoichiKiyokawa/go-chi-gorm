package seeds

import (
	"fmt"
	"go-mux-gorm/entity"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	var users []entity.User
	for i := 0; i < 10; i++ {
		users = append(users, entity.User{Name: fmt.Sprintf("user%d", i), Email: fmt.Sprintf("user%d@example.com", i), Password: "foobar"})
	}
	db.Create(users)
}
