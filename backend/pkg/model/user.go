package model

type User struct {
	Id          int    `json:"-" gorm:"primaryKey;unique; not null"`
	Username    string `json:"username" binding:"required" gorm:"unique; not null; type:varchar(255)"`
	Email       string `json:"email" binding:"required" gorm:"not null; type:varchar(255)"`
	PhoneNumber string `json:"phone_number" binding:"required" gorm:"not null; type:varchar(15)"`
	Password    string `json:"password" binding:"required" gorm:"not null; type:varchar(255)"`
}

type Role struct {
	Username string `gorm:"uniqueIndex:not null; type:varchar(255)"`
	Role     string `gorm:"uniqueIndex:not null; type:varchar(255)"`
}

type UserLoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
