package usecase

import (
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"github.com/kzuabe/todolist-go-api/internal/repository"
)

type UserUseCase struct {
	Repository repository.UserRepositoryInterface
}

type UserUseCaseInterface interface {
	FetchByID(int) (entity.User, error)
	Create(entity.User) (entity.User, error)
}

func (useCase *UserUseCase) FetchByID(id int) (entity.User, error) {
	return useCase.Repository.FetchByID(id)
}

func (useCase *UserUseCase) Create(user entity.User) (entity.User, error) {
	return useCase.Repository.Create(user)
}
