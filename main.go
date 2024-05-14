package main

import (
	"class-room-repair/initapp"
	"class-room-repair/logger"
	"class-room-repair/server"
	"embed"
	"go.uber.org/zap"

	_ "github.com/swaggo/gin-swagger"
)

//go:embed resources/config/application.yml
var yamlFile embed.FS

//go:embed resources/config/zaplogger.yml
var zapYamlFile embed.FS

func main() {
	initapp.Init()
	r := server.NewRouter()
	err := r.Run(":8080")
	if err != nil {
		logger.ZapLogger.Fatal("start server failed", zap.Error(err))
	}
}
