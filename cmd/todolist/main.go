package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/config"
	"github.com/kzuabe/todolist-go-api/internal/controller"
	"github.com/kzuabe/todolist-go-api/internal/repository"
	"github.com/kzuabe/todolist-go-api/internal/router"
	"github.com/kzuabe/todolist-go-api/internal/usecase"
	"github.com/kzuabe/todolist-go-api/pkg/middleware"
)

// @title                       TodoList API
// @version                     1.0
// @description                 TODOリストアプリのRESTfulAPI（Go実装）
// @securitydefinitions.apikey  TokenAuth
// @in                          header
// @name                        Authorization
func main() {
	// 環境ごとのセットアップ
	if config.API_ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := repository.NewDB()
	if err != nil {
		panic("failed to connect database")
	}
	repository := repository.NewTaskRepository(db)
	usecase := usecase.NewTaskUseCase(repository)
	controller := controller.NewTaskController(usecase)
	faMiddleware, err := middleware.NewFirebaseAuthMiddleware()
	if err != nil {
		log.Fatalf("error initializing firebase admin: %v\n", err)
	}

	r := router.NewRouter(controller, faMiddleware)

	r.Run(":8080")
}
