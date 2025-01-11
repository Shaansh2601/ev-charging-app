package model

import "time"

type Station struct {
	Id        int       `json:"-" gorm:"primaryKey"`
	UserId    int       `json:"user_id" gorm:""`
	Address   string    `json:"address" gorm:""`
	CreatedAt time.Time `json:"created_at" gorm:""`
	UpdatedAt time.Time `json:"updated_at" gorm:""`
}

type ChargingPoint struct { //Have to think how to do it better
	Id            int       `json:"-" gorm:"primaryKey"`
	IsActive      bool      `json:"is_active" gorm:""`
	ChargingSpeed int       `json:"charging_speed" gorm:""`
	SupportType   string    `json:"support_type" gorm:""`
	UpdatedAt     time.Time `json:"updated_at" gorm:""`
}
