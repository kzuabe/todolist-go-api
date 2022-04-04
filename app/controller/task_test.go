package controller

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewTaskController(t *testing.T) {
	type args struct {
		useCase TaskUseCaseInterface
	}
	tests := []struct {
		name string
		args args
		want *TaskController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskController(tt.args.useCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskController_Get(t *testing.T) {
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
			controller.Get(tt.args.c)
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
