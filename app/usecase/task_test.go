package usecase

import (
	"reflect"
	"testing"

	"github.com/kzuabe/todolist-go-api/app/model"
)

func TestTaskUseCase_Fetch(t *testing.T) {
	type fields struct {
		Repository taskRepository
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
			useCase := &TaskUseCase{
				Repository: tt.fields.Repository,
			}
			got, err := useCase.Fetch(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskUseCase.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskUseCase.Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskUseCase_FetchByID(t *testing.T) {
	type fields struct {
		Repository taskRepository
	}
	type args struct {
		id     string
		userID string
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
			useCae := &TaskUseCase{
				Repository: tt.fields.Repository,
			}
			got, err := useCae.FetchByID(tt.args.id, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskUseCase.FetchByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskUseCase.FetchByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskUseCase_Create(t *testing.T) {
	type fields struct {
		Repository taskRepository
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
			useCase := &TaskUseCase{
				Repository: tt.fields.Repository,
			}
			got, err := useCase.Create(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskUseCase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskUseCase_Update(t *testing.T) {
	type fields struct {
		Repository taskRepository
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
			useCase := &TaskUseCase{
				Repository: tt.fields.Repository,
			}
			got, err := useCase.Update(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskUseCase.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskUseCase.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskUseCase_Delete(t *testing.T) {
	type fields struct {
		Repository taskRepository
	}
	type args struct {
		id     string
		userID string
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
			useCase := &TaskUseCase{
				Repository: tt.fields.Repository,
			}
			if err := useCase.Delete(tt.args.id, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("TaskUseCase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verifyUser(tt.args.task, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("verifyUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
