package controller

import (
	"bytes"
	"errors"
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
		/*
			Get Tests
		*/
		{
			name: "Get: リクエストに対して正常なレスポンスを返す",
			args: args{
				method: http.MethodGet,
				target: "/",
				token: &auth.Token{
					UID: "user_1",
				},
			},
			mocks: []mock{
				{
					funcName: "Fetch",
					args: []any{
						model.TaskFetchParam{UserID: "user_1"},
					},
					returnArgs: []any{
						[]model.Task{
							{
								ID:          "task_1",
								UserID:      "user_1",
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
				body: "../../test/testdata/res_tasks.json",
			},
		},
		{
			name: "Get: リクエストに対して正常なレスポンスを返す（クエリパラメータあり）",
			args: args{
				method: http.MethodGet,
				target: "/?status=0",
				token: &auth.Token{
					UID: "user_1",
				},
			},
			mocks: []mock{
				{
					funcName: "Fetch",
					args: []any{
						model.TaskFetchParam{UserID: "user_1", Status: newPoint(0)},
					},
					returnArgs: []any{
						[]model.Task{
							{
								ID:          "task_1",
								UserID:      "user_1",
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
				body: "../../test/testdata/res_tasks.json",
			},
		},
		{
			name: "Get: リクエストに対してエラーレスポンスを返す",
			args: args{
				method: http.MethodGet,
				target: "/",
				token: &auth.Token{
					UID: "user_1",
				},
			},
			mocks: []mock{
				{
					funcName: "Fetch",
					args: []any{
						model.TaskFetchParam{UserID: "user_1"},
					},
					returnArgs: []any{
						[]model.Task{},
						errors.New("テストエラーメッセージ"),
					},
				},
			},
			want: want{
				code: 500,
				body: "../../test/testdata/res_error.json",
			},
		},

		/*
			GetByID Tests
		*/
		{
			name: "GetByID: リクエストに対して正常なレスポンスを返す",
			args: args{
				method: http.MethodGet,
				target: "/task_1",
				token: &auth.Token{
					UID: "user_1",
				},
			},
			mocks: []mock{
				{
					funcName: "FetchByID",
					args:     []any{"task_1", "user_1"},
					returnArgs: []any{
						model.Task{
							ID:          "task_1",
							UserID:      "user_1",
							Title:       "テストタスク1",
							Description: "テストタスク説明1",
							Status:      0,
						},
						nil,
					},
				},
			},
			want: want{
				code: 200,
				body: "../../test/testdata/res_task.json",
			},
		},
		{
			name: "GetByID: リクエストに対してエラーレスポンスを返す",
			args: args{
				method: http.MethodGet,
				target: "/task_1",
				token: &auth.Token{
					UID: "user_1",
				},
			},
			mocks: []mock{
				{
					funcName: "FetchByID",
					args:     []any{"task_1", "user_1"},
					returnArgs: []any{
						model.Task{},
						errors.New("テストエラーメッセージ"),
					},
				},
			},
			want: want{
				code: 500,
				body: "../../test/testdata/res_error.json",
			},
		},

		/*
			Post Tests
		*/
		{
			name: "Post: リクエストに対して正常なレスポンスを返す",
			args: args{
				method: http.MethodPost,
				target: "/",
				body:   "../../test/testdata/req_task.json",
				token: &auth.Token{
					UID: "user_1",
				},
			},
			mocks: []mock{
				{
					funcName: "Create",
					args: []any{
						model.Task{
							UserID:      "user_1",
							Title:       "テストタスク1",
							Description: "テストタスク説明1",
							Status:      0,
						},
					},
					returnArgs: []any{
						model.Task{
							ID:          "task_1",
							UserID:      "user_1",
							Title:       "テストタスク1",
							Description: "テストタスク説明1",
							Status:      0,
						},
						nil,
					},
				},
			},
			want: want{
				code: 201,
				body: "../../test/testdata/res_task.json",
			},
		},
		{
			name: "Post: リクエストに対してエラーレスポンスを返す",
			args: args{
				method: http.MethodPost,
				target: "/",
				body:   "../../test/testdata/req_task.json",
				token: &auth.Token{
					UID: "user_1",
				},
			},
			mocks: []mock{
				{
					funcName: "Create",
					args: []any{
						model.Task{
							UserID:      "user_1",
							Title:       "テストタスク1",
							Description: "テストタスク説明1",
							Status:      0,
						},
					},
					returnArgs: []any{
						model.Task{},
						errors.New("テストエラーメッセージ"),
					},
				},
			},
			want: want{
				code: 500,
				body: "../../test/testdata/res_error.json",
			},
		},

		/*
			Put Tests
		*/
		{
			name: "Put: リクエストに対して正常なレスポンスを返す",
			args: args{
				method: http.MethodPut,
				target: "/task_1",
				body:   "../../test/testdata/req_task.json",
				token: &auth.Token{
					UID: "user_1",
				},
			},
			mocks: []mock{
				{
					funcName: "Update",
					args: []any{
						model.Task{
							ID:          "task_1",
							UserID:      "user_1",
							Title:       "テストタスク1",
							Description: "テストタスク説明1",
							Status:      0,
						},
					},
					returnArgs: []any{
						model.Task{
							ID:          "task_1",
							UserID:      "user_1",
							Title:       "テストタスク1",
							Description: "テストタスク説明1",
							Status:      0,
						},
						nil,
					},
				},
			},
			want: want{
				code: 200,
				body: "../../test/testdata/res_task.json",
			},
		},
		{
			name: "Put: リクエストに対してエラーレスポンスを返す",
			args: args{
				method: http.MethodPut,
				target: "/task_1",
				body:   "../../test/testdata/req_task.json",
				token: &auth.Token{
					UID: "user_1",
				},
			},
			mocks: []mock{
				{
					funcName: "Update",
					args: []any{
						model.Task{
							ID:          "task_1",
							UserID:      "user_1",
							Title:       "テストタスク1",
							Description: "テストタスク説明1",
							Status:      0,
						},
					},
					returnArgs: []any{
						model.Task{},
						errors.New("テストエラーメッセージ"),
					},
				},
			},
			want: want{
				code: 500,
				body: "../../test/testdata/res_error.json",
			},
		},

		/*
			Delete Tests
		*/
		{
			name: "Delete: リクエストに対して正常なレスポンスを返す",
			args: args{
				method: http.MethodDelete,
				target: "/task_1",
				token: &auth.Token{
					UID: "user_1",
				},
			},
			mocks: []mock{
				{
					funcName:   "Delete",
					args:       []any{"task_1", "user_1"},
					returnArgs: []any{nil},
				},
			},
			want: want{
				code: 204,
			},
		},
		{
			name: "Delete: リクエストに対してエラーレスポンスを返す",
			args: args{
				method: http.MethodDelete,
				target: "/task_1",
				token: &auth.Token{
					UID: "user_1",
				},
			},
			mocks: []mock{
				{
					funcName:   "Delete",
					args:       []any{"task_1", "user_1"},
					returnArgs: []any{errors.New("テストエラーメッセージ")},
				},
			},
			want: want{
				code: 500,
				body: "../../test/testdata/res_error.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ルーターのセットアップ
			mockUseCase := new(mocks.TaskUseCaseInterface)
			for _, mock := range tt.mocks {
				mockUseCase.On(mock.funcName, mock.args...).Return(mock.returnArgs...)
			}
			controller := NewTaskController(mockUseCase)
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
			if body := tt.want.body; body != "" {
				wantBody, err := ioutil.ReadFile(tt.want.body)
				if err != nil {
					t.Errorf("Response Body File not found: %v", tt.want.body)
				}
				assert.Equal(t, string(wantBody), w.Body.String())
			}
			mockUseCase.AssertExpectations(t)
		})
	}
}

func newPoint[V any](v V) *V {
	return &v
}
