package usecase

import (
	"net/http"

	"github.com/kzuabe/todolist-go-api/app/model"
)

type TaskRepositoryInterface interface {
	Fetch(model.TaskFetchParam) ([]model.Task, error)
	FetchByID(string) (model.Task, error)
	Create(model.Task) (model.Task, error)
	Update(model.Task) (model.Task, error)
	Delete(string) error
}

type TaskUseCase struct {
	Repository TaskRepositoryInterface
}

func NewTaskUseCase(repository TaskRepositoryInterface) *TaskUseCase {
	return &TaskUseCase{Repository: repository}
}

func (useCase *TaskUseCase) Fetch(params model.TaskFetchParam) ([]model.Task, error) {
	if params.UserID == "" { // ユーザーIDの指定必須
		err := &model.Error{StatusCode: http.StatusInternalServerError, Message: "ユーザー情報が取得できませんでした"}
		return []model.Task{}, err
	}

	return useCase.Repository.Fetch(params)
}

func (useCase *TaskUseCase) FetchByID(id string, userID string) (model.Task, error) {
	task, err := useCase.Repository.FetchByID(id)
	if err != nil {
		return model.Task{}, err
	}

	// ユーザーの検証
	if err := verifyUser(task, userID); err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (useCase *TaskUseCase) Create(task model.Task) (model.Task, error) {
	return useCase.Repository.Create(task)
}

func (useCase *TaskUseCase) Update(task model.Task) (model.Task, error) {
	// ユーザーの検証
	registered, err := useCase.Repository.FetchByID(task.ID)
	if err != nil {
		return model.Task{}, err
	}
	if err = verifyUser(registered, task.UserID); err != nil {
		return model.Task{}, err
	}

	return useCase.Repository.Update(task)
}

func (useCase *TaskUseCase) Delete(id string, userID string) error {
	// ユーザーの検証
	registered, err := useCase.Repository.FetchByID(id)
	if err != nil {
		return err
	}
	if err := verifyUser(registered, userID); err != nil {
		return err
	}

	return useCase.Repository.Delete(id)
}

// ユーザーIDをもとにタスクの権限があるかを検証する（権限が無い場合はエラーを返す）
func verifyUser(task model.Task, userID string) error {
	if task.UserID != userID {
		err := &model.Error{StatusCode: http.StatusForbidden, Message: "許可されていないユーザーからのリクエストです"}
		return err
	}
	return nil
}
