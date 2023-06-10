package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-gin-frame/internal/router"
)

func main() {
	engine := gin.Default()

	gin.SetMode(gin.DebugMode)
	// Switch to "release" mode in production
	// gin.SetMode(gin.ReleaseMode)

	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 设置 API 路由
	router.SetApiRouter(engine)
	// 设置 Socket 路由
	router.SetSocketRouter(engine)

	engine.Run(":8080")
}
