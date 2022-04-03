package repository

import (
	"reflect"
	"testing"

	"github.com/kzuabe/todolist-go-api/app/model"
	"gorm.io/gorm"
)

func TestTaskRepository_Fetch(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		params model.TaskFetchParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &TaskRepository{
				DB: tt.fields.DB,
			}
			got, err := repository.Fetch(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskRepository.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskRepository.Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRepository_FetchByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &TaskRepository{
				DB: tt.fields.DB,
			}
			got, err := repository.FetchByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskRepository.FetchByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskRepository.FetchByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRepository_Create(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		task model.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &TaskRepository{
				DB: tt.fields.DB,
			}
			got, err := repository.Create(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRepository_Update(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		task model.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &TaskRepository{
				DB: tt.fields.DB,
			}
			got, err := repository.Update(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRepository_Delete(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &TaskRepository{
				DB: tt.fields.DB,
			}
			if err := repository.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TaskRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
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
		// TODO: Add test cases.
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
			if got := task.toModel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.toModel() = %v, want %v", got, tt.want)
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toDBTask(tt.args.task); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toDBTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
