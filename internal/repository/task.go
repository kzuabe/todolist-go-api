package repository

import (
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

type TaskRepositoryInterface interface {
	Fetch(entity.TaskFetchParam) ([]entity.Task, error)
	Create(entity.Task) (entity.Task, error)
}

type Task struct {
	gorm.Model
	UUID        string
	UserID      string
	Title       string
	Description string
	Status      int
}

func (repository *TaskRepository) Fetch(params entity.TaskFetchParam) ([]entity.Task, error) {
	tx := repository.DB.Session(&gorm.Session{})
	if status := params.Status; status != nil {
		tx = tx.Where("status = ?", status)
	}

	tasks := []entity.Task{}
	result := tx.Find(&tasks, "user_id = ?", params.UserID)
	return tasks, result.Error
}

func (repository *TaskRepository) Create(task entity.Task) (entity.Task, error) {
	t := toDBTask(task)
	result := repository.DB.Create(&t)
	created := toEntityTask(t)
	return created, result.Error
}

func toDBTask(task entity.Task) Task {
	t := Task{
		UUID:        task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
	return t
}

func toEntityTask(task Task) entity.Task {
	t := entity.Task{
		ID:          task.UUID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
	return t
}
