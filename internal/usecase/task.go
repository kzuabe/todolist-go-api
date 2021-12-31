package usecase

import (
	"strings"

	"github.com/google/uuid"
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"github.com/kzuabe/todolist-go-api/internal/repository"
)

type TaskUseCase struct {
	Repository repository.TaskRepositoryInterface
}

type TaskUseCaseInterface interface {
	Fetch(entity.TaskFetchParam) ([]entity.Task, error)
	Create(entity.Task) (entity.Task, error)
	Update(entity.Task) (entity.Task, error)
}

func (useCase *TaskUseCase) Fetch(params entity.TaskFetchParam) ([]entity.Task, error) {
	return useCase.Repository.Fetch(params)
}

func (useCase *TaskUseCase) Create(task entity.Task) (entity.Task, error) {
	uuid := strings.ReplaceAll(uuid.NewString(), "-", "") // UUIDの生成（ハイフン除去済み）
	task.ID = uuid
	return useCase.Repository.Create(task)
}

func (useCase *TaskUseCase) Update(task entity.Task) (entity.Task, error) {
	return useCase.Repository.Update(task)
}
