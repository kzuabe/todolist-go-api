package controller

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestTaskController_Get(t *testing.T) {
	type fields struct {
		UseCase taskUseCase
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
		UseCase taskUseCase
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
		UseCase taskUseCase
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
		UseCase taskUseCase
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
		UseCase taskUseCase
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
