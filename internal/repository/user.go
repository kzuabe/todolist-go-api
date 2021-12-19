package repository

import (
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	FetchByID(int) (entity.User, error)
	Create(entity.User) (entity.User, error)
}

type User struct {
	gorm.Model
	Name string
}

func (repository *UserRepository) FetchByID(id int) (entity.User, error) {
	return entity.User{ID: id}, nil
}

func (repository *UserRepository) Create(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}
