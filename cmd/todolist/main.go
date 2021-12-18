package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DSN") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.User{}, &entity.Task{})

	router := gin.Default()

	router.Run("localhost:8080")
}
