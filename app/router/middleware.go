package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/app/model"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		if e, ok := err.Err.(*model.Error); ok {
			c.IndentedJSON(e.StatusCode, e)
			return
		} else {
			e = &model.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
			c.IndentedJSON(e.StatusCode, e)
			return
		}
	}
}
