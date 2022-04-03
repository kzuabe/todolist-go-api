//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kzuabe/todolist-go-api/app/controller"
	"github.com/kzuabe/todolist-go-api/app/repository"
	"github.com/kzuabe/todolist-go-api/app/router"
	"github.com/kzuabe/todolist-go-api/app/usecase"
)

func initializeRouter() (*gin.Engine, error) {
	wire.Build(
		router.NewRouter,
		controller.NewTaskController,
		usecase.NewTaskUseCase,
		repository.NewTaskRepository,
		repository.NewDB,

		wire.Bind(new(usecase.TaskRepositoryInterface), new(*repository.TaskRepository)),
		wire.Bind(new(controller.TaskUseCaseInterface), new(*usecase.TaskUseCase)),
	)
	return &gin.Engine{}, nil
}
