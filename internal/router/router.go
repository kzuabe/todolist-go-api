package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/controller"
	"github.com/kzuabe/todolist-go-api/pkg/middleware"
)

type Handler struct {
	UserController         controller.UserControllerInterface
	FirebaseAuthMiddleware middleware.FirebaseAuthMiddlewareInterface
}

func NewRouter(h Handler) *gin.Engine {
	router := gin.Default()

	router.Use(h.FirebaseAuthMiddleware.MiddlewareFunc())
	v1 := router.Group("/v1")
	{
		v1.GET("/users/:id", h.UserController.GetByID)
		v1.POST("/users", h.UserController.Post)
	}

	return router
}
