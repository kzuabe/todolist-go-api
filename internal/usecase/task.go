package usecase

import (
	"net/http"

	"github.com/kzuabe/todolist-go-api/internal/model"
)

type TaskRepositoryInterface interface {
	Fetch(model.TaskFetchParam) ([]model.Task, error)
	FetchByID(string) (model.Task, error)
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
	task, err := useCae.Repository.FetchByID(id)
	if err != nil {
		return model.Task{}, err
	}

	// リクエストユーザーとタスクのユーザーが異なる場合はエラーを返す
	if task.UserID != userID {
		err = &model.Error{StatusCode: http.StatusForbidden, Message: "許可されていないユーザーからのリクエストです"}
		return model.Task{}, err
	}

	return task, nil
}

func (useCase *TaskUseCase) Create(task model.Task) (model.Task, error) {
	return useCase.Repository.Create(task)
}

func (useCase *TaskUseCase) Update(task model.Task) (model.Task, error) {
	return useCase.Repository.Update(task)
}

func (useCase *TaskUseCase) Delete(id string, userID string) error {
	return useCase.Repository.Delete(id, userID)
}
