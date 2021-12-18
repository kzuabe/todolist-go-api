package usecase

import (
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"github.com/kzuabe/todolist-go-api/internal/repository"
)

type UserUseCase struct {
	Repository repository.UserRepositoryInterface
}

type UserUseCaseInterface interface {
	FetchByID(uint) (entity.User, error)
}

func (useCase *UserUseCase) FetchByID(id uint) (entity.User, error) {
	return useCase.Repository.FetchByID(id)
}
