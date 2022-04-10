package controller

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/ginauth"
	"github.com/kzuabe/todolist-go-api/app/model"
	mocks "github.com/kzuabe/todolist-go-api/test/mocks/app/controller"
	"github.com/stretchr/testify/assert"
)

func newTestTaskRouter(controller *TaskController, token ginauth.FirebaseAuthToken) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set(ginauth.FirebaseAuthTokenKey, token)
	})
	r.GET("/", controller.Get)
	r.GET("/:id", controller.GetByID)
	r.POST("/", controller.Post)
	r.PUT("/:id", controller.Put)
	r.DELETE("/:id", controller.Delete)
	return r
}

func TestTaskController_Get(t *testing.T) {
	type args struct {
		r     *http.Request
		token ginauth.FirebaseAuthToken
	}
	type expects struct {
		argParam model.TaskFetchParam
		tasks    []model.Task
		err      error
	}
	type want struct {
		code int
		body string
	}
	tests := []struct {
		name    string
		args    args
		expects expects
		want    want
	}{
		{
			name: "リクエストに対して正常なレスポンスを返す",
			args: args{
				r: httptest.NewRequest("GET", "/", nil),
			},
			expects: expects{
				tasks: []model.Task{
					{
						ID:          "testtaskid1",
						UserID:      "testuserid1",
						Title:       "テストタスク1",
						Description: "テストタスク説明1",
						Status:      0,
					},
				},
			},
			want: want{
				code: 200,
				body: "../../test/testdata/res_get_tasks.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ルーターのセットアップ
			useCase := new(mocks.TaskUseCaseInterface)
			useCase.On("Fetch", tt.expects.argParam).Return(tt.expects.tasks, tt.expects.err)
			controller := NewTaskController(useCase)
			router := newTestTaskRouter(controller, tt.args.token)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, tt.args.r)

			assert.Equal(t, tt.want.code, w.Code)
			wantBody, err := ioutil.ReadFile(tt.want.body)
			if err != nil {
				t.Errorf("Body File not found: %v", tt.want.body)
			}
			assert.Equal(t, wantBody, w.Body.String())
		})
	}
}

func TestTaskController_GetByID(t *testing.T) {
	type fields struct {
		UseCase TaskUseCaseInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &TaskController{
				UseCase: tt.fields.UseCase,
			}
			controller.GetByID(tt.args.c)
		})
	}
}

func TestTaskController_Post(t *testing.T) {
	type fields struct {
		UseCase TaskUseCaseInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &TaskController{
				UseCase: tt.fields.UseCase,
			}
			controller.Post(tt.args.c)
		})
	}
}

func TestTaskController_Put(t *testing.T) {
	type fields struct {
		UseCase TaskUseCaseInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &TaskController{
				UseCase: tt.fields.UseCase,
			}
			controller.Put(tt.args.c)
		})
	}
}

func TestTaskController_Delete(t *testing.T) {
	type fields struct {
		UseCase TaskUseCaseInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &TaskController{
				UseCase: tt.fields.UseCase,
			}
			controller.Delete(tt.args.c)
		})
	}
}
