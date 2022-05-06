package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kzuabe/ginauth"
	"github.com/kzuabe/todolist-go-api/app/controller"
	_ "github.com/kzuabe/todolist-go-api/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter(tc *controller.TaskController) *gin.Engine {
	router := gin.Default()

	router.Use(controller.ErrorHandler())
	router.Use(cors.New(corsConfig()))

	v1 := router.Group("/v1")

	provider := ginauth.NewFirebaseAuthProvider()
	v1.Use(ginauth.NewAuthorizer(provider))
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

func corsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	return config
}
