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

	router.GET("/", h.UserController.GetByID)

	return router
}
