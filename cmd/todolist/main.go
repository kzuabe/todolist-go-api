package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kzuabe/todolist-go-api/internal/config"
)

// @title                       TodoList API
// @version                     1.0
// @description                 TODOリストアプリのRESTfulAPI（Go実装）
// @securitydefinitions.apikey  TokenAuth
// @in                          header
// @name                        Authorization
func main() {
	// 本番・開発環境のセットアップ
	if config.API_ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r, err := initializeRouter()
	if err != nil {
		log.Fatalf("error initializing: %v\n", err)
	}

	r.Run(":8080")
}
