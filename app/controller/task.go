package controller

import (
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/ginauth"
	"github.com/kzuabe/todolist-go-api/app/model"
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
// @Param        status  query     int  false  "タスクステータス 0: 未着手 1: 完了"  Enums(0, 1)
// @Success      200     {object}  []model.Task
// @Failure      401     "Unauthorized"
// @Failure      500     {object}  model.Error
// @Security     TokenAuth
// @Router       /v1/tasks [get]
func (controller *TaskController) Get(c *gin.Context) {
	token := c.MustGet(ginauth.FirebaseAuthTokenKey).(*auth.Token)

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
// @Param        id   path  string  true  "タスクID"
// @Success      200      {object}  model.Task
// @Failure      401      "Unauthorized"
// @Failure      403      {object}  model.Error
// @Failure      404      {object}  model.Error
// @Failure      500      {object}  model.Error
// @Security     TokenAuth
// @Router       /v1/tasks/{id} [get]
func (controller *TaskController) GetByID(c *gin.Context) {
	token := c.MustGet(ginauth.FirebaseAuthTokenKey).(*auth.Token)

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
// @Param        payload  body      model.Task  true  "登録タスク内容（id / user_idの値は自動セット）"
// @Success      201      {object}  model.Task
// @Failure      401      "Unauthorized"
// @Failure      500      {object}  model.Error
// @Security     TokenAuth
// @Router       /v1/tasks [post]
func (controller *TaskController) Post(c *gin.Context) {
	token := c.MustGet(ginauth.FirebaseAuthTokenKey).(*auth.Token)

	task := model.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		e := &model.Error{StatusCode: http.StatusBadRequest, Message: err.Error()}
		c.Error(e)
		return
	}
	task.UserID = token.UID // タスクのUserIDはトークンの値をセットし直す

	created, err := controller.UseCase.Create(task)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusCreated, created)
}

// Put godoc
// @Summary      タスク更新
// @Description  ユーザのタスクを更新する
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id       path      string      true  "タスクID"
// @Param        payload  body      model.Task  true  "登録タスク内容（id / user_idの値は自動セット）"
// @Success      200  {object}  model.Task
// @Failure      401  "Unauthorized"
// @Failure      403  {object}  model.Error
// @Failure      404  {object}  model.Error
// @Failure      500  {object}  model.Error
// @Security     TokenAuth
// @Router       /v1/tasks/{id} [put]
func (controller *TaskController) Put(c *gin.Context) {
	token, _ := c.MustGet(ginauth.FirebaseAuthTokenKey).(*auth.Token)

	task := model.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		e := &model.Error{StatusCode: http.StatusBadRequest, Message: err.Error()}
		c.Error(e)
		return
	}

	// IDとUserIDは他パラメータからセットし直す
	task.ID = c.Param("id")
	task.UserID = token.UID

	updated, err := controller.UseCase.Update(task)
	if err != nil {
		c.Error(err)
		return
	}
	c.IndentedJSON(http.StatusOK, updated)
}

// Delete godoc
// @Summary      タスク削除
// @Description  ユーザのタスクを削除する
// @Tags         task
// @Param        id   path      string  true  "タスクID"
// @Success      204  "No Content"
// @Failure      401  "Unauthorized"
// @Failure      403  {object}  model.Error
// @Failure      404  {object}  model.Error
// @Failure      500  {object}  model.Error
// @Security     TokenAuth
// @Router       /v1/tasks/{id} [delete]
func (controller *TaskController) Delete(c *gin.Context) {
	token := c.MustGet(ginauth.FirebaseAuthTokenKey).(*auth.Token)

	id := c.Param("id")
	userID := token.UID

	err := controller.UseCase.Delete(id, userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
