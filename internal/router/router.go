package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kzuabe/todolist-go-api/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type taskController interface {
	Get(c *gin.Context)
	GetByID(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}

type firebaseAuthMiddleware interface {
	MiddlewareFunc() gin.HandlerFunc
}

type Handler struct {
	TaskController         taskController
	FirebaseAuthMiddleware firebaseAuthMiddleware
}

func NewHandler(tc taskController, fam firebaseAuthMiddleware) Handler {
	return Handler{TaskController: tc, FirebaseAuthMiddleware: fam}
}

func NewRouter(h Handler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.Use(h.FirebaseAuthMiddleware.MiddlewareFunc())
	{
		v1.GET("/tasks", h.TaskController.Get)
		v1.GET("tasks/:id", h.TaskController.GetByID)
		v1.POST("/tasks", h.TaskController.Post)
		v1.PUT("/tasks/:id", h.TaskController.Put)
		v1.DELETE("/tasks/:id", h.TaskController.Delete)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
