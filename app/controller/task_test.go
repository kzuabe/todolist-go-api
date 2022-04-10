package controller

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/ginauth"
	"github.com/kzuabe/todolist-go-api/app/model"
	mocks "github.com/kzuabe/todolist-go-api/test/mocks/app/controller"
	"github.com/stretchr/testify/assert"
)

func newTestTaskRouter(controller *TaskController, token *auth.Token) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set(ginauth.FirebaseAuthTokenKey, token)
	})
	r.Use(ErrorHandler())
	r.GET("/", controller.Get)
	r.GET("/:id", controller.GetByID)
	r.POST("/", controller.Post)
	r.PUT("/:id", controller.Put)
	r.DELETE("/:id", controller.Delete)
	return r
}

func TestTaskController(t *testing.T) {
	type args struct {
		method string
		target string
		body   string
		token  *auth.Token
	}
	type mock struct {
		funcName   string
		args       []any
		returnArgs []any
	}
	type want struct {
		code int
		body string
	}
	tests := []struct {
		name  string
		args  args
		mocks []mock
		want  want
	}{
		{
			name: "リクエストに対して正常なレスポンスを返す",
			args: args{
				method: http.MethodGet,
				target: "/",
				token: &auth.Token{
					UID: "testuserid1",
				},
			},
			mocks: []mock{
				{
					funcName: "Fetch",
					args: []any{
						model.TaskFetchParam{UserID: "testuserid1"},
					},
					returnArgs: []any{
						[]model.Task{
							{
								ID:          "testtaskid1",
								UserID:      "testuserid1",
								Title:       "テストタスク1",
								Description: "テストタスク説明1",
								Status:      0,
							},
						},
						nil,
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
			for _, mock := range tt.mocks {
				useCase.On(mock.funcName, mock.args...).Return(mock.returnArgs...)
			}
			controller := NewTaskController(useCase)
			router := newTestTaskRouter(controller, tt.args.token)

			// リクエストの作成
			var reqBody io.Reader
			if body := tt.args.body; body != "" {
				b, err := ioutil.ReadFile(body)
				if err != nil {
					t.Errorf("Request Body File not found: %v", tt.want.body)
				}
				reqBody = bytes.NewBuffer(b)
			}
			r := httptest.NewRequest(tt.args.method, tt.args.target, reqBody)

			// リクエスト実行
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)

			// テスト実行
			assert.Equal(t, tt.want.code, w.Code)
			wantBody, err := ioutil.ReadFile(tt.want.body)
			if err != nil {
				t.Errorf("Response Body File not found: %v", tt.want.body)
			}
			assert.Equal(t, string(wantBody), w.Body.String())
		})
	}
}
