package service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type ChatService struct {
	repo repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) CreatePost(post model.Post) (int, error) {
	return s.repo.CreatePost(post)
}

func (s *ChatService) GetAllPosts() ([]model.Post, error) {
	return s.repo.GetAllPosts()
}

func (s *ChatService) CreateComment(comment model.Comment) (int, error) {
	return s.repo.CreateComment(comment)
}

func (s *ChatService) GetAllComments(postId int) ([]model.Comment, error) {
	return s.repo.GetAllComments(postId)
}
