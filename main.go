package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-gin-frame/internal/router"
)

func main() {
	engine := gin.Default()

	gin.SetMode(gin.DebugMode)
	// Switch to "release" mode in production
	// gin.SetMode(gin.ReleaseMode)

	// 跨域
	engine.Use(cors.Default())

	// 设置 API 路由
	router.SetApiRouter(engine)
	// 设置 Socket 路由
	router.SetSocketRouter(engine)

	engine.Run(":8080")
}
