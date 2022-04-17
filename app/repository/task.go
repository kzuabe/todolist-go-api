package repository

import (
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/kzuabe/todolist-go-api/app/model"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (repository *TaskRepository) Fetch(params model.TaskFetchParam) ([]model.Task, error) {
	tx := repository.DB.Session(&gorm.Session{})
	if userID := params.UserID; userID != "" {
		tx = tx.Where("user_id = ?", userID)
	}
	if status := params.Status; status != nil {
		tx = tx.Where("status = ?", status)
	}

	dbTasks := []Task{}
	result := tx.Find(&dbTasks)

	if err := result.Error; err != nil {
		return []model.Task{}, wrapError(err)
	}

	tasks := make([]model.Task, len(dbTasks))
	for i, t := range dbTasks {
		tasks[i] = t.toModel()
	}
	return tasks, nil
}

func (repository *TaskRepository) FetchByID(id string) (model.Task, error) {
	dbTask := Task{}

	result := repository.DB.First(&dbTask, "uuid = ?", id)

	if err := result.Error; err != nil {
		return model.Task{}, wrapError(err)
	}

	fetched := dbTask.toModel()
	return fetched, nil
}

func (repository *TaskRepository) Create(task model.Task) (model.Task, error) {
	dbTask := toDBTask(task)
	id := strings.ReplaceAll(uuid.NewString(), "-", "") // UUIDの生成（ハイフン除去済み）
	dbTask.UUID = id                                    // 生成時はUUIDを自動でセット

	result := repository.DB.Create(&dbTask)

	if err := result.Error; err != nil {
		return model.Task{}, wrapError(err)
	}

	created := dbTask.toModel()
	return created, result.Error
}

func (repository *TaskRepository) Update(task model.Task) (model.Task, error) {
	dbTask := Task{}

	result := repository.DB.First(&dbTask, "uuid = ?", task.ID)
	if err := result.Error; err != nil {
		return model.Task{}, wrapError(err)
	}

	dbTask.updateFromModel(task) // 更新

	result = repository.DB.Save(&dbTask)
	if err := result.Error; err != nil {
		return model.Task{}, wrapError(err)
	}

	updated := dbTask.toModel()
	return updated, nil
}

func (repository *TaskRepository) Delete(id string) error {
	dbTask := Task{}

	result := repository.DB.First(&dbTask, "uuid = ?", id)
	if err := result.Error; err != nil {
		return wrapError(err)
	}

	result = repository.DB.Delete(&dbTask)
	if err := result.Error; err != nil {
		return wrapError(err)
	}
	return nil
}

type Task struct {
	gorm.Model
	UUID        string
	UserID      string
	Title       string
	Description string
	Status      int
}

func (task *Task) toModel() model.Task {
	t := model.Task{
		ID:          task.UUID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
	return t
}

func toDBTask(task model.Task) Task {
	t := Task{
		UUID:        task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
	return t
}

// WARNING: カラム追加時はここに更新するフィールドの処理が必要
func (task *Task) updateFromModel(t model.Task) {
	task.UserID = t.UserID
	task.Title = t.Title
	task.Description = t.Description
	task.Status = t.Status
}

func wrapError(err error) error {
	var wrapped error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		wrapped = &model.Error{StatusCode: http.StatusNotFound, Message: "データが存在しません"}
		return wrapped
	}
	wrapped = &model.Error{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	return wrapped
}
