package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/controller"
)

type Handler struct {
	UserController controller.UserControllerInterface
}

func NewRouter(h Handler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/users/:id", h.UserController.GetByID)
		v1.POST("/users", h.UserController.Post)
	}

	return router
}
