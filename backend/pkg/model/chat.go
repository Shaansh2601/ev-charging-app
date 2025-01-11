package model

import "time"

type Post struct {
	Id          int       `json:"-" gorm:"primaryKey"`
	Username    string    `json:"-" gorm:""`
	Title       string    `json:"title" gorm:""`
	Description string    `json:"description" gorm:""`
	CreatedAt   time.Time `json:"created_at" gorm:""`
}

type Comment struct {
	Id        int       `json:"-" gorm:"uniqueIndex:message_id;"`
	PostID    int       `json:"post_id,omitempty" gorm:"uniqueIndex:message_id;"`
	Username  string    `json:"-" gorm:""`
	Message   string    `json:"message" gorm:""`
	CreatedAt time.Time `json:"created_at" gorm:""`
}
