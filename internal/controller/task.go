package controller

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/model"
	"github.com/kzuabe/todolist-go-api/pkg/middleware"
)

type TaskUseCaseInterface interface {
	Fetch(model.TaskFetchParam) ([]model.Task, error)
	FetchByID(string, string) (model.Task, error)
	Create(model.Task) (model.Task, error)
	Update(model.Task) (model.Task, error)
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
// @Param        status   query     int     false  "タスクステータス 0: 未着手 1: 完了"  Enums(0, 1)
// @Success      200      {object}  []model.Task
// @Security     TokenAuth
// @Router       /v1/tasks [get]
func (controller *TaskController) Get(c *gin.Context) {
	token := c.MustGet(middleware.CONTEXT_TOKEN_KEY).(*auth.Token)

	params := model.TaskFetchParam{}
	if err := c.ShouldBindQuery(&params); err != nil {
		e := &model.Error{StatusCode: http.StatusBadRequest, Message: err.Error()}
		c.Error(e)
		return
	}
	params.UserID = token.UID

	tasks, err := controller.UseCase.Fetch(params)
	if err != nil {
		c.Error(err)
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

// GetByID godoc
// @Summary      タスク取得（1件）
// @Description  ユーザのタスクを1件取得する
// @Tags         task
// @Produce      json
// @Param        id   path      string  true  "タスクID"
// @Success      200      {object}  model.Task
// @Security     TokenAuth
// @Router       /v1/tasks/{id} [get]
func (controller *TaskController) GetByID(c *gin.Context) {
	token := c.MustGet(middleware.CONTEXT_TOKEN_KEY).(*auth.Token)

	id := c.Param("id")
	userID := token.UID

	task, err := controller.UseCase.FetchByID(id, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

// Post godoc
// @Summary      タスク追加
// @Description  ユーザのタスクを追加する
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        payload  body      model.Task  true  "Payload Description"
// @Success      200      {object}  model.Task
// @Security     TokenAuth
// @Router       /v1/tasks [post]
func (controller *TaskController) Post(c *gin.Context) {
	token, _ := c.MustGet(middleware.CONTEXT_TOKEN_KEY).(*auth.Token)

	task := model.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if task.UserID != token.UID {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "ユーザーIDが認証情報と一致しません"})
		return
	}

	created, _ := controller.UseCase.Create(task)

	c.IndentedJSON(http.StatusCreated, created)
}

// Put godoc
// @Summary      タスク更新
// @Description  ユーザのタスクを更新する
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        payload  body      model.Task  true  "Payload Description"
// @Success      200  {object}  model.Task
// @Security     TokenAuth
// @Router       /v1/tasks/{id} [put]
func (controller *TaskController) Put(c *gin.Context) {
	token, _ := c.MustGet(middleware.CONTEXT_TOKEN_KEY).(*auth.Token)

	task := model.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if task.ID != c.Param("id") {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "タスクIDが一致しません"})
		return
	}
	if task.UserID != token.UID {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "ユーザーIDが認証情報と一致しません"})
		return
	}

	updated, _ := controller.UseCase.Update(task)
	c.IndentedJSON(http.StatusCreated, updated)
}

// Delete godoc
// @Summary      タスク削除
// @Description  ユーザのタスクを削除する
// @Tags         task
// @Produce      json
// @Param        id   path      string  true  "タスクID"
// @Success      200  {object}  model.Task
// @Security     TokenAuth
// @Router       /v1/tasks/{id} [delete]
func (controller *TaskController) Delete(c *gin.Context) {
	token, _ := c.MustGet(middleware.CONTEXT_TOKEN_KEY).(*auth.Token)

	id := c.Param("id")
	userID := token.UID

	_ = controller.UseCase.Delete(id, userID)
	c.Status(http.StatusNoContent)
}
