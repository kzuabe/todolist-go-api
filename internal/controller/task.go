package controller

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/entity"
	"github.com/kzuabe/todolist-go-api/pkg/middleware"
)

type TaskUseCaseInterface interface {
	Fetch(entity.TaskFetchParam) ([]entity.Task, error)
	FetchByID(string, string) (entity.Task, error)
	Create(entity.Task) (entity.Task, error)
	Update(entity.Task) (entity.Task, error)
	Delete(string, string) error
}

type TaskController struct {
	UseCase TaskUseCaseInterface
}

func NewTaskController(useCase TaskUseCaseInterface) *TaskController {
	return &TaskController{UseCase: useCase}
}

// Get godoc
// @Summary      タスク取得
// @Description  ユーザのタスクを複数件取得する
// @Tags         task
// @Produce      json
// @Param        status  query     int  false  "タスクステータス 0: 未着手 1: 完了"  Enums(0, 1)
// @Success      200     {object}  []entity.Task
// @Security     TokenAuth
// @Router       /v1/tasks [get]
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

// GetByID godoc
// @Summary      タスク取得（1件）
// @Description  ユーザのタスクを1件取得する
// @Tags         task
// @Produce      json
// @Param        id   path      string  true  "タスクID"
// @Success      200      {object}  entity.Task
// @Security     TokenAuth
// @Router       /v1/tasks/{id} [get]
func (controller *TaskController) GetByID(c *gin.Context) {
	token, _ := c.MustGet(middleware.CONTEXT_TOKEN_KEY).(*auth.Token)
	id := c.Param("id")
	task, _ := controller.UseCase.FetchByID(id, token.UID)
	c.IndentedJSON(http.StatusOK, task)
}

// Post godoc
// @Summary      タスク追加
// @Description  ユーザのタスクを追加する
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        payload  body      entity.Task  true  "Payload Description"
// @Success      200      {object}  entity.Task
// @Security     TokenAuth
// @Router       /v1/tasks [post]
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

// Put godoc
// @Summary      タスク更新
// @Description  ユーザのタスクを更新する
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        payload  body      entity.Task  true  "Payload Description"
// @Success      200  {object}  entity.Task
// @Security     TokenAuth
// @Router       /v1/tasks/{id} [put]
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

// Delete godoc
// @Summary      タスク削除
// @Description  ユーザのタスクを削除する
// @Tags         task
// @Produce      json
// @Param        id   path      string  true  "タスクID"
// @Success      200  {object}  entity.Task
// @Security     TokenAuth
// @Router       /v1/tasks/{id} [delete]
func (controller *TaskController) Delete(c *gin.Context) {
	token, _ := c.MustGet(middleware.CONTEXT_TOKEN_KEY).(*auth.Token)

	id := c.Param("id")
	userID := token.UID

	_ = controller.UseCase.Delete(id, userID)
	c.Status(http.StatusNoContent)
}
