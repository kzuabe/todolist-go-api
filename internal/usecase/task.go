package usecase

import (
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"github.com/kzuabe/todolist-go-api/internal/repository"
)

type TaskUseCase struct {
	Repository repository.TaskRepositoryInterface
}

type TaskUseCaseInterface interface {
	Fetch() ([]entity.Task, error)
}

func (useCase *TaskUseCase) Fetch() ([]entity.Task, error) {
	return useCase.Repository.Fetch()
}
