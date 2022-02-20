package repository

import (
	"github.com/kzuabe/todolist-go-api/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() (*gorm.DB, error) {
	dsn := config.DSN + "?charset=utf8mb4&parseTime=True&loc=Local"
	// ログレベルを環境ごとに設定
	var logLevel = logger.Info
	if config.API_ENV == "production" {
		logLevel = logger.Error
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Task{})
	return db, nil
}
