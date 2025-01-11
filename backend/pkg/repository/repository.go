package repository

import (
	"backend/pkg/model"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User, role string) (int, error)
	GetUser(username, password string) ([]string, error)
}

type Chat interface {
	CreatePost(post model.Post) (int, error)
	GetAllPosts() ([]model.Post, error)
	CreateComment(message model.Comment) (int, error)
	GetAllComments(postId int) ([]model.Comment, error)
}

type Repository struct {
	Authorization
	Chat
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Chat:          NewChatDB(db),
	}
}
