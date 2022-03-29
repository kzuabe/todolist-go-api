package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kzuabe/todolist-go-api/docs"
	"github.com/kzuabe/todolist-go-api/internal/controller"
	"github.com/kzuabe/todolist-go-api/pkg/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter(tc *controller.TaskController, fc middleware.Client) *gin.Engine {
	router := gin.Default()

	router.Use(ErrorHandler())

	v1 := router.Group("/v1")
	v1.Use(middleware.NewAuthorizer(fc))
	{
		v1.GET("/tasks", tc.Get)
		v1.GET("tasks/:id", tc.GetByID)
		v1.POST("/tasks", tc.Post)
		v1.PUT("/tasks/:id", tc.Put)
		v1.DELETE("/tasks/:id", tc.Delete)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
