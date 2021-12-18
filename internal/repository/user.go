package repository

import (
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	FetchByID(uint) (entity.User, error)
}

type User struct {
	gorm.Model
	Name string
}

func (repository *UserRepository) FetchByID(id uint) (entity.User, error) {
	return entity.User{}, nil
}
