package controller

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"github.com/kzuabe/todolist-go-api/internal/usecase"
	"github.com/kzuabe/todolist-go-api/pkg/middleware"
)

type TaskController struct {
	UseCase usecase.TaskUseCaseInterface
}

type TaskControllerInterface interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}

func (controller *TaskController) Get(c *gin.Context) {
	token, _ := c.MustGet(middleware.CONTEXT_TOKEN_KEY).(*auth.Token)

	params := entity.TaskFetchParam{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	params.UserID = token.UID

	tasks, _ := controller.UseCase.Fetch(params)
	c.IndentedJSON(http.StatusOK, tasks)
}

func (controller *TaskController) Post(c *gin.Context) {
	token, _ := c.MustGet(middleware.CONTEXT_TOKEN_KEY).(*auth.Token)

	task := entity.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	task.UserID = token.UID

	created, _ := controller.UseCase.Create(task)

	c.IndentedJSON(http.StatusCreated, created)
}

func (controller *TaskController) Put(c *gin.Context) {
	token, _ := c.MustGet(middleware.CONTEXT_TOKEN_KEY).(*auth.Token)

	task := entity.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	task.ID = c.Param("id")
	task.UserID = token.UID

	updated, _ := controller.UseCase.Update(task)
	c.IndentedJSON(http.StatusCreated, updated)
}

func (controller *TaskController) Delete(c *gin.Context) {
	token, _ := c.MustGet(middleware.CONTEXT_TOKEN_KEY).(*auth.Token)

	id := c.Param("id")
	userID := token.UID

	_ = controller.UseCase.Delete(id, userID)
	c.Status(http.StatusNoContent)
}
