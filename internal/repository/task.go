package repository

import (
	"github.com/kzuabe/todolist-go-api/internal/model"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

type Task struct {
	gorm.Model
	UUID        string
	UserID      string
	Title       string
	Description string
	Status      int
}

func (repository *TaskRepository) Fetch(params model.TaskFetchParam) ([]model.Task, error) {
	tx := repository.DB.Session(&gorm.Session{})
	if status := params.Status; status != nil {
		tx = tx.Where("status = ?", status)
	}

	dbTasks := []Task{}
	result := tx.Find(&dbTasks, "user_id = ?", params.UserID)

	tasks := make([]model.Task, len(dbTasks))
	for i, t := range dbTasks {
		tasks[i] = t.toModel()
	}
	return tasks, result.Error
}

func (repository *TaskRepository) FetchByID(id string, userID string) (model.Task, error) {
	t := Task{}

	result := repository.DB.First(&t, "uuid = ?", id)
	if t.UserID != userID { // TODO: エラーハンドリング
		return model.Task{}, nil
	}
	return t.toModel(), result.Error
}

func (repository *TaskRepository) Create(task model.Task) (model.Task, error) {
	t := toDBTask(task)
	result := repository.DB.Create(&t)
	created := t.toModel()
	return created, result.Error
}

func (repository *TaskRepository) Update(task model.Task) (model.Task, error) {
	t := Task{}
	_ = repository.DB.First(&t, "uuid = ?", task.ID) // TODO: エラーハンドリング
	if t.UserID != task.UserID {
		return model.Task{}, nil
	}

	// 更新内容
	t.Title = task.Title
	t.Description = task.Description
	t.Status = task.Status

	result := repository.DB.Save(&t)
	return t.toModel(), result.Error
}

func (repository *TaskRepository) Delete(id string, userID string) error {
	result := repository.DB.Delete(&Task{}, "uuid = ? AND user_id = ?", id, userID)
	return result.Error
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
