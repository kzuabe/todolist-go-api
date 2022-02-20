package usecase

import (
	"strings"

	"github.com/google/uuid"
	"github.com/kzuabe/todolist-go-api/internal/entity"
)

type TaskRepositoryInterface interface {
	Fetch(entity.TaskFetchParam) ([]entity.Task, error)
	FetchByID(string, string) (entity.Task, error)
	Create(entity.Task) (entity.Task, error)
	Update(entity.Task) (entity.Task, error)
	Delete(string, string) error
}

type TaskUseCase struct {
	Repository TaskRepositoryInterface
}

func NewTaskUseCase(repository TaskRepositoryInterface) *TaskUseCase {
	return &TaskUseCase{Repository: repository}
}

func (useCase *TaskUseCase) Fetch(params entity.TaskFetchParam) ([]entity.Task, error) {
	return useCase.Repository.Fetch(params)
}

func (useCae *TaskUseCase) FetchByID(id string, userID string) (entity.Task, error) {
	return useCae.Repository.FetchByID(id, userID)
}

func (useCase *TaskUseCase) Create(task entity.Task) (entity.Task, error) {
	uuid := strings.ReplaceAll(uuid.NewString(), "-", "") // UUIDの生成（ハイフン除去済み）
	task.ID = uuid
	return useCase.Repository.Create(task)
}

func (useCase *TaskUseCase) Update(task entity.Task) (entity.Task, error) {
	return useCase.Repository.Update(task)
}

func (useCase *TaskUseCase) Delete(id string, userID string) error {
	return useCase.Repository.Delete(id, userID)
}
