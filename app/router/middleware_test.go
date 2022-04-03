package router

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestErrorHandler(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
