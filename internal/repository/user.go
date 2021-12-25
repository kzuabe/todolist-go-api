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
	u := User{}
	result := repository.DB.First(&u, id)
	return toEntityUser(u), result.Error
}

func (repository *UserRepository) Create(user entity.User) (entity.User, error) {
	u := toRepositoryUser(user)
	result := repository.DB.Create(&u)
	return toEntityUser(u), result.Error
}

func toEntityUser(user User) entity.User {
	u := entity.User{
		ID:   int(user.ID),
		Name: user.Name,
	}
	return u
}

func toRepositoryUser(user entity.User) User {
	u := User{
		Name: user.Name,
	}
	return u
}
