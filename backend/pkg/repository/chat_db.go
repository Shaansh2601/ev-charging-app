package repository

import (
	"backend/pkg/model"

	"gorm.io/gorm"
)

type ChatDB struct {
	db *gorm.DB
}

func NewChatDB(db *gorm.DB) *ChatDB {
	return &ChatDB{db: db}
}

func (r *ChatDB) CreatePost(post model.Post) (int, error) {
	tx := r.db.Begin()

	err := tx.Create(&post).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return post.Id, tx.Commit().Error
}

func (r *ChatDB) GetAllPosts() ([]model.Post, error) {
	var posts []model.Post
	tx := r.db.Begin()
	result := r.db.Find(&posts)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	return posts, tx.Commit().Error
}

func (r *ChatDB) CreateComment(comment model.Comment) (int, error) {
	tx := r.db.Begin()

	err := tx.Create(&comment).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return comment.Id, tx.Commit().Error
}

func (r *ChatDB) GetAllComments(postId int) ([]model.Comment, error) {
	var comments []model.Comment

	tx := r.db.Begin()
	result := tx.Where("post_id = ?", postId).Find(&comments)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	return comments, tx.Commit().Error
}
