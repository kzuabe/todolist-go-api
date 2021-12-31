package repository

import (
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

type TaskRepositoryInterface interface {
	Fetch() ([]entity.Task, error)
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

func (repository *TaskRepository) Fetch() ([]entity.Task, error) {
	return []entity.Task{}, nil
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
