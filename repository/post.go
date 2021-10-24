package repository

import (
	"go-chi-gorm/entity"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) FindOne(id int) (entity.Post, error) {
	var post entity.Post
	result := r.db.Find(&post, id)
	return post, result.Error
}

func (r *PostRepository) FindAll() ([]entity.Post, error) {
	var posts []entity.Post
	result := r.db.Find(&posts)
	return posts, result.Error
}

func (r *PostRepository) FindAllByUserId(userID int) ([]entity.Post, error) {
	var posts []entity.Post
	result := r.db.Joins("User").Where("posts.user_id = ?", userID).Find(&posts)
	return posts, result.Error
}

func (r *PostRepository) Create(user entity.User) (entity.User, error) {
	result := r.db.Create(&user)
	return user, result.Error
}
