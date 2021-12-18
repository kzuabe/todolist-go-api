package main

import (
	"os"

	"github.com/kzuabe/todolist-go-api/internal/controller"
	"github.com/kzuabe/todolist-go-api/internal/repository"
	"github.com/kzuabe/todolist-go-api/internal/router"
	"github.com/kzuabe/todolist-go-api/internal/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DSN") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&repository.User{})

	h := router.Handler{
		UserController: &controller.UserController{
			UseCase: &usecase.UserUseCase{
				Repository: &repository.UserRepository{
					DB: db,
				},
			},
		},
	}
	r := router.NewRouter(h)

	r.Run(":8080")
}
