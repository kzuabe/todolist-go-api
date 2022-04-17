package usecase

import (
	"testing"

	"github.com/kzuabe/todolist-go-api/app/model"
	mocks "github.com/kzuabe/todolist-go-api/test/mocks/app/usecase"
	"github.com/stretchr/testify/assert"
)

type mock struct {
	funcName   string
	args       []any
	returnArgs []any
}

func TestTaskUseCase_Fetch(t *testing.T) {
	type args struct {
		params model.TaskFetchParam
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Task
		wantErr bool
	}{
		{
			name: "UserIDの値が空文字だった場合にエラーを返す",
			args: args{
				params: model.TaskFetchParam{UserID: ""},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := new(mocks.TaskRepositoryInterface)
			useCase := NewTaskUseCase(mockRepository)

			got, err := useCase.Fetch(tt.args.params)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})
	}
}

func TestTaskUseCase_FetchByID(t *testing.T) {
	type args struct {
		id     string
		userID string
	}
	tests := []struct {
		name    string
		args    args
		mocks   []mock
		want    model.Task
		wantErr bool
	}{
		{
			name: "取得したタスクのUserIDが入力と一致しなかった場合にエラーを返す",
			args: args{
				id:     "task_id1",
				userID: "user_id2",
			},
			mocks: []mock{
				{
					funcName: "FetchByID",
					args:     []any{"task_id1"},
					returnArgs: []any{
						model.Task{ID: "task_id1", UserID: "user_id1"},
						nil,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := new(mocks.TaskRepositoryInterface)
			for _, mock := range tt.mocks {
				mockRepository.On(mock.funcName, mock.args...).Return(mock.returnArgs...)
			}
			useCase := NewTaskUseCase(mockRepository)
			got, err := useCase.FetchByID(tt.args.id, tt.args.userID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})
	}
}

func TestTaskUseCase_Update(t *testing.T) {
	type args struct {
		task model.Task
	}
	tests := []struct {
		name    string
		args    args
		mocks   []mock
		want    model.Task
		wantErr bool
	}{
		{
			name: "更新したいタスクのUserIDが入力と一致しなかった場合にエラーを返す",
			args: args{
				task: model.Task{ID: "task_id1", UserID: "user_id2"},
			},
			mocks: []mock{
				{
					funcName: "FetchByID",
					args:     []any{"task_id1"},
					returnArgs: []any{
						model.Task{ID: "task_id1", UserID: "user_id1"},
						nil,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := new(mocks.TaskRepositoryInterface)
			for _, mock := range tt.mocks {
				mockRepository.On(mock.funcName, mock.args...).Return(mock.returnArgs...)
			}
			useCase := NewTaskUseCase(mockRepository)
			got, err := useCase.Update(tt.args.task)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})
	}
}

func TestTaskUseCase_Delete(t *testing.T) {
	type args struct {
		id     string
		userID string
	}
	tests := []struct {
		name    string
		args    args
		mocks   []mock
		wantErr bool
	}{
		{
			name: "削除したいタスクのUserIDが入力と一致しなかった場合にエラーを返す",
			args: args{
				id:     "task_id1",
				userID: "user_id2",
			},
			mocks: []mock{
				{
					funcName: "FetchByID",
					args:     []any{"task_id1"},
					returnArgs: []any{
						model.Task{ID: "task_id1", UserID: "user_id1"},
						nil,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := new(mocks.TaskRepositoryInterface)
			for _, mock := range tt.mocks {
				mockRepository.On(mock.funcName, mock.args...).Return(mock.returnArgs...)
			}
			useCase := NewTaskUseCase(mockRepository)
			err := useCase.Delete(tt.args.id, tt.args.userID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}

func Test_verifyUser(t *testing.T) {
	type args struct {
		task   model.Task
		userID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "UserIDの値が一致する場合はエラーを返さない",
			args: args{
				task:   model.Task{UserID: "user_id1"},
				userID: "user_id1",
			},
			wantErr: false,
		},
		{
			name: "UserIDの値が一致しない場合はエラーを返す",
			args: args{
				task:   model.Task{UserID: "user_id1"},
				userID: "user_id2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := verifyUser(tt.args.task, tt.args.userID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
