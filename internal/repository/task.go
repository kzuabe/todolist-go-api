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
}

type Task struct {
	gorm.Model
	UserID      string
	Title       string
	Description string
	Status      int
}

func (repository *TaskRepository) Fetch() ([]entity.Task, error) {
	return []entity.Task{}, nil
}
