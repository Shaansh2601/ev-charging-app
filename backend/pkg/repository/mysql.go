package repository

import (
	"backend/pkg/model"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewMySQLDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Role{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Post{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Comment{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
