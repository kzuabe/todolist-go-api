package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/app/config"
)

// @title                       TodoList API
// @version                     1.0
// @description                 TODOリストアプリのRESTfulAPI（Go実装）
// @securitydefinitions.apikey  TokenAuth
// @in                          header
// @name                        Authorization
func main() {
	gin.SetMode(config.GinMode)

	r, err := initializeRouter()
	if err != nil {
		log.Fatalf("error initializing: %v\n", err)
	}

	r.Run(":8080")
}
