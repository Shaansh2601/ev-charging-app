package service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User, role string) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, []string, error)
}

type Chat interface {
	CreatePost(post model.Post) (int, error)
	GetAllPosts() ([]model.Post, error)
	CreateComment(message model.Comment) (int, error)
	GetAllComments(postId int) ([]model.Comment, error)
}

type Service struct {
	Authorization
	Chat
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Chat:          NewChatService(repos.Chat),
	}
}
