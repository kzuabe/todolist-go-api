package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"github.com/kzuabe/todolist-go-api/internal/usecase"
)

type UserController struct {
	UseCase usecase.UserUseCaseInterface
}

type UserControllerInterface interface {
	GetByID(c *gin.Context)
	Post(c *gin.Context)
}

func (controller *UserController) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := controller.UseCase.FetchByID(id)
	c.IndentedJSON(http.StatusOK, user)
}

func (controller *UserController) Post(c *gin.Context) {
	user, _ := controller.UseCase.Create(entity.User{Name: "テストユーザ"})
	c.IndentedJSON(http.StatusCreated, user)
}
