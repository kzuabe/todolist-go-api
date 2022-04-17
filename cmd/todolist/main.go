package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/app/config"
	"github.com/kzuabe/todolist-go-api/app/controller"
	"github.com/kzuabe/todolist-go-api/app/repository"
	"github.com/kzuabe/todolist-go-api/app/router"
	"github.com/kzuabe/todolist-go-api/app/usecase"
)

// @title                       TodoList API
// @version                     1.0
// @description                 TODOリストアプリのRESTfulAPI（Go実装）
// @securitydefinitions.apikey  TokenAuth
// @in                          header
// @name                        Authorization
func main() {
	gin.SetMode(config.GinMode)

	r, err := initRouter()
	if err != nil {
		log.Fatalf("error initializing: %v\n", err)
	}

	r.Run(":8080")
}

// Dependency Injection
func initRouter() (*gin.Engine, error) {
	db := repository.NewDB()
	taskRepository := repository.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	taskController := controller.NewTaskController(taskUseCase)
	router := router.NewRouter(taskController)
	return router, nil
}
