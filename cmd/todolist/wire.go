//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kzuabe/todolist-go-api/internal/controller"
	"github.com/kzuabe/todolist-go-api/internal/repository"
	"github.com/kzuabe/todolist-go-api/internal/router"
	"github.com/kzuabe/todolist-go-api/internal/usecase"
	"github.com/kzuabe/todolist-go-api/pkg/middleware"
)

func initializeRouter() (*gin.Engine, error) {
	wire.Build(
		router.NewRouter,
		controller.NewTaskController,
		usecase.NewTaskUseCase,
		repository.NewTaskRepository,
		repository.NewDB,
		middleware.NewFirebaseAuthMiddleware,

		wire.Bind(new(usecase.TaskRepositoryInterface), new(*repository.TaskRepository)),
		wire.Bind(new(controller.TaskUseCaseInterface), new(*usecase.TaskUseCase)),
	)
	return &gin.Engine{}, nil
}
