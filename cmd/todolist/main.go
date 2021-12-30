package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/kzuabe/todolist-go-api/internal/controller"
	"github.com/kzuabe/todolist-go-api/internal/repository"
	"github.com/kzuabe/todolist-go-api/internal/router"
	"github.com/kzuabe/todolist-go-api/internal/usecase"
	"github.com/kzuabe/todolist-go-api/pkg/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// DB セットアップ
	dsn := os.Getenv("DSN") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&repository.Task{})

	// Firebase Admin セットアップ
	firebaseApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	firebaseAuthClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	h := router.Handler{
		TaskController: &controller.TaskController{
			UseCase: &usecase.TaskUseCase{
				Repository: &repository.TaskRepository{
					DB: db,
				},
			},
		},
		FirebaseAuthMiddleware: &middleware.FirebaseAuthMiddleware{
			Client: firebaseAuthClient,
		},
	}
	r := router.NewRouter(h)

	r.Run(":8080")
}
