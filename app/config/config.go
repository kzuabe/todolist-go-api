package config

import (
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

// 設定値
var GinMode string
var GormLogLevel logger.LogLevel
var DSN string

func init() {
	// 環境ごとの設定値をセット（デフォルトは本番設定）
	switch env := os.Getenv("API_ENV"); env {
	case "develop":
		GinMode = gin.DebugMode
		GormLogLevel = logger.Info
	default:
		GinMode = gin.ReleaseMode
		GormLogLevel = logger.Error
	}
	DSN = os.Getenv("DSN")
}
