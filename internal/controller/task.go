package controller

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"github.com/kzuabe/todolist-go-api/internal/usecase"
)

type TaskController struct {
	UseCase usecase.TaskUseCaseInterface
}

type TaskControllerInterface interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
}

func (controller *TaskController) Get(c *gin.Context) {
	tasks, _ := controller.UseCase.Fetch()
	c.IndentedJSON(http.StatusOK, tasks)
}

func (controller *TaskController) Post(c *gin.Context) {
	token, _ := c.MustGet("token").(*auth.Token)

	task := entity.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	task.UserID = token.UID

	created, _ := controller.UseCase.Create(task)

	c.IndentedJSON(http.StatusCreated, created)
}
