package router

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewRouter(h Handler) *gin.Engine {
	router := gin.Default()

	return router
}
