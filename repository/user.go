package repository

import (
	"go-mux-gorm/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FindOne(id int) (entity.User, error) {
	var user entity.User
	result := u.db.Find(&user, id)
	return user, result.Error
}

func (u *UserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	result := u.db.Find(&users)
	return users, result.Error
}

func (u *UserRepository) Create(user entity.User) (entity.User, error) {
	result := u.db.Create(&user)
	return user, result.Error
}
