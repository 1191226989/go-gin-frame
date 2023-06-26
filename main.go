package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-gin-frame/assets"
	"go-gin-frame/global"
	"go-gin-frame/internal/router"
)

func main() {
	global.Init()

	engine := gin.Default()

	gin.SetMode(gin.DebugMode)
	// Switch to "release" mode in production
	// gin.SetMode(gin.ReleaseMode)

	// 初始化默认静态资源 http://127.0.0.1:8080/assets/static
	engine.StaticFS("assets", http.FS(assets.Static))

	// 跨域
	engine.Use(cors.Default())

	// 设置 API 路由
	router.SetApiRouter(engine)
	// 设置 Socket 路由
	router.SetSocketRouter(engine)

	engine.Run(":8080")
}
