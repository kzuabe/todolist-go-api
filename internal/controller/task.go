package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/usecase"
)

type TaskController struct {
	UseCase usecase.TaskUseCaseInterface
}

type TaskControllerInterface interface {
	Get(c *gin.Context)
}

func (controller *TaskController) Get(c *gin.Context) {
	tasks, _ := controller.UseCase.Fetch()
	c.IndentedJSON(http.StatusOK, tasks)
}
