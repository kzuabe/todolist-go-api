package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/usecase"
)

type UserController struct {
	UseCase usecase.UserUseCaseInterface
}

type UserControllerInterface interface {
	GetByID(c *gin.Context)
}

func (controller *UserController) GetByID(c *gin.Context) {
	user, _ := controller.UseCase.FetchByID(1)
	c.IndentedJSON(http.StatusOK, user)
}
