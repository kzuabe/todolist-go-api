package usecase

import (
	"strings"

	"github.com/google/uuid"
	"github.com/kzuabe/todolist-go-api/internal/model"
)

type TaskRepositoryInterface interface {
	Fetch(model.TaskFetchParam) ([]model.Task, error)
	FetchByID(string, string) (model.Task, error)
	Create(model.Task) (model.Task, error)
	Update(model.Task) (model.Task, error)
	Delete(string, string) error
}

type TaskUseCase struct {
	Repository TaskRepositoryInterface
}

func NewTaskUseCase(repository TaskRepositoryInterface) *TaskUseCase {
	return &TaskUseCase{Repository: repository}
}

func (useCase *TaskUseCase) Fetch(params model.TaskFetchParam) ([]model.Task, error) {
	return useCase.Repository.Fetch(params)
}

func (useCae *TaskUseCase) FetchByID(id string, userID string) (model.Task, error) {
	return useCae.Repository.FetchByID(id, userID)
}

func (useCase *TaskUseCase) Create(task model.Task) (model.Task, error) {
	uuid := strings.ReplaceAll(uuid.NewString(), "-", "") // UUIDの生成（ハイフン除去済み）
	task.ID = uuid
	return useCase.Repository.Create(task)
}

func (useCase *TaskUseCase) Update(task model.Task) (model.Task, error) {
	return useCase.Repository.Update(task)
}

func (useCase *TaskUseCase) Delete(id string, userID string) error {
	return useCase.Repository.Delete(id, userID)
}
