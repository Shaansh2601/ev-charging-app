package model

type Transaction struct {
	Id   int `json:"-" gorm:"primaryKey;unique; not null"`
	Cost int `json:"cost" binding:"required" gorm:"not null;"`
	//Consumption int `json:"consumption,omitempty" gorm:"not null;"`
}
