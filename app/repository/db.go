package repository

import (
	"github.com/kzuabe/todolist-go-api/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	dsn := config.DSN + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(config.GormLogLevel),
	})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Task{})
	return db
}
