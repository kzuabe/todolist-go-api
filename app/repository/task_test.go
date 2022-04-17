package repository

import (
	"testing"

	"github.com/kzuabe/todolist-go-api/app/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTestTasks() []Task {
	tasks := []Task{
		{
			UUID:        "task_id1",
			UserID:      "user_id1",
			Title:       "テストタスク1",
			Description: "テストタスク説明1",
			Status:      0,
		},
		{
			UUID:        "task_id2",
			UserID:      "user_id2",
			Title:       "テストタスク2",
			Description: "テストタスク説明2",
			Status:      1,
		},
	}
	return tasks
}

func newTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Task{})
	tasks := newTestTasks()
	db.Create(&tasks)
	return db, nil
}

func TestTaskRepository_Fetch(t *testing.T) {
	db, err := newTestDB()
	if err != nil {
		t.Errorf("error: %s", err.Error())
		return
	}
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
			name: "全てのタスクを取得する場合",
			args: args{},
			want: []model.Task{
				{
					ID:          "task_id1",
					UserID:      "user_id1",
					Title:       "テストタスク1",
					Description: "テストタスク説明1",
					Status:      0,
				},
				{
					ID:          "task_id2",
					UserID:      "user_id2",
					Title:       "テストタスク2",
					Description: "テストタスク説明2",
					Status:      1,
				},
			},
		},
		{
			name: "UserIDの指定がある場合に対象ユーザーのタスクのみ返す",
			args: args{
				params: model.TaskFetchParam{UserID: "user_id1"},
			},
			want: []model.Task{
				{
					ID:          "task_id1",
					UserID:      "user_id1",
					Title:       "テストタスク1",
					Description: "テストタスク説明1",
					Status:      0,
				},
			},
		},
		{
			name: "Statusの指定がある場合にStatusの一致するタスクのみ返す",
			args: args{
				params: model.TaskFetchParam{Status: newPoint(1)},
			},
			want: []model.Task{
				{
					ID:          "task_id2",
					UserID:      "user_id2",
					Title:       "テストタスク2",
					Description: "テストタスク説明2",
					Status:      1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := NewTaskRepository(db)
			got, err := repository.Fetch(tt.args.params)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})
	}
}

func TestTaskRepository_FetchByID(t *testing.T) {
	db, err := newTestDB()
	if err != nil {
		t.Errorf("error: %s", err.Error())
		return
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Task
		wantErr bool
	}{
		{
			name: "IDを元にタスクを取得する",
			args: args{id: "task_id1"},
			want: model.Task{
				ID:          "task_id1",
				UserID:      "user_id1",
				Title:       "テストタスク1",
				Description: "テストタスク説明1",
				Status:      0,
			},
		},
		{
			name:    "存在しないIDを指定した場合にエラーを返す",
			args:    args{id: "task_id3"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := NewTaskRepository(db)
			got, err := repository.FetchByID(tt.args.id)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})
	}
}

func TestTaskRepository_Create(t *testing.T) {
	db, err := newTestDB()
	if err != nil {
		t.Errorf("error: %s", err.Error())
		return
	}
	type args struct {
		task model.Task
	}
	tests := []struct {
		name    string
		args    args
		want    model.Task
		wantErr bool
	}{
		{
			name: "入力を元にタスクを登録する",
			args: args{
				task: model.Task{
					UserID:      "user_id3",
					Title:       "テストタスク3",
					Description: "テストタスク説明3",
					Status:      0,
				},
			},
			want: model.Task{
				UserID:      "user_id3",
				Title:       "テストタスク3",
				Description: "テストタスク説明3",
				Status:      0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := NewTaskRepository(db)
			got, err := repository.Create(tt.args.task)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			tt.want.ID = got.ID // IDがランダム生成のため
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})
	}
}

func TestTaskRepository_Update(t *testing.T) {
	db, err := newTestDB()
	if err != nil {
		t.Errorf("error: %s", err.Error())
		return
	}
	type args struct {
		task model.Task
	}
	tests := []struct {
		name    string
		args    args
		want    model.Task
		wantErr bool
	}{
		{
			name: "入力を元にタスクを更新する",
			args: args{
				task: model.Task{
					ID:          "task_id1",
					UserID:      "user_id3",
					Title:       "テストタスク3",
					Description: "テストタスク説明3",
					Status:      1,
				},
			},
			want: model.Task{
				ID:          "task_id1",
				UserID:      "user_id3",
				Title:       "テストタスク3",
				Description: "テストタスク説明3",
				Status:      1,
			},
		},
		{
			name: "存在しないIDを指定した場合にエラーを返す",
			args: args{
				task: model.Task{ID: "task_id3"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := NewTaskRepository(db)
			got, err := repository.Update(tt.args.task)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})
	}
}

func TestTaskRepository_Delete(t *testing.T) {
	db, err := newTestDB()
	if err != nil {
		t.Errorf("error: %s", err.Error())
		return
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "入力したIDのタスクを削除する",
			args: args{id: "task_id1"},
		},
		{
			name:    "存在しないIDを指定した場合にエラーを返す",
			args:    args{id: "task_id3"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := NewTaskRepository(db)
			err := repository.Delete(tt.args.id)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}

func TestTask_toModel(t *testing.T) {
	type fields struct {
		Model       gorm.Model
		UUID        string
		UserID      string
		Title       string
		Description string
		Status      int
	}
	tests := []struct {
		name   string
		fields fields
		want   model.Task
	}{
		{
			name: "repository.Taskをmodel.Taskに変換する",
			fields: fields{
				UUID:        "task_id1",
				UserID:      "user_id1",
				Title:       "テストタスク1",
				Description: "テストタスク説明1",
				Status:      0,
			},
			want: model.Task{
				ID:          "task_id1",
				UserID:      "user_id1",
				Title:       "テストタスク1",
				Description: "テストタスク説明1",
				Status:      0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &Task{
				Model:       tt.fields.Model,
				UUID:        tt.fields.UUID,
				UserID:      tt.fields.UserID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
			}
			got := task.toModel()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_toDBTask(t *testing.T) {
	type args struct {
		task model.Task
	}
	tests := []struct {
		name string
		args args
		want Task
	}{
		{
			name: "model.Taskをrepository.Taskに変換する",
			args: args{
				task: model.Task{
					ID:          "task_id1",
					UserID:      "user_id1",
					Title:       "テストタスク1",
					Description: "テストタスク説明1",
					Status:      0,
				},
			},
			want: Task{
				UUID:        "task_id1",
				UserID:      "user_id1",
				Title:       "テストタスク1",
				Description: "テストタスク説明1",
				Status:      0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toDBTask(tt.args.task)
			assert.Equal(t, tt.want, got)
		})
	}
}

func newPoint[V any](v V) *V {
	return &v
}
