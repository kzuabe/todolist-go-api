package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/controller"
	"github.com/kzuabe/todolist-go-api/pkg/middleware"
)

type Handler struct {
	TaskController         controller.TaskControllerInterface
	FirebaseAuthMiddleware middleware.FirebaseAuthMiddlewareInterface
}

func NewRouter(h Handler) *gin.Engine {
	router := gin.Default()

	router.Use(h.FirebaseAuthMiddleware.MiddlewareFunc())
	v1 := router.Group("/v1")
	{
		v1.GET("/tasks", h.TaskController.Get)
		v1.GET("tasks/:id", h.TaskController.GetByID)
		v1.POST("/tasks", h.TaskController.Post)
		v1.PUT("/tasks/:id", h.TaskController.Put)
		v1.DELETE("/tasks/:id", h.TaskController.Delete)
	}

	return router
}
