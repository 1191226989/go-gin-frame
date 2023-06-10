package router

import (
	"go-gin-frame/internal/api/article"
	"go-gin-frame/internal/api/socket"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	// Simple group: v1
	group := router.Group("/api")
	{
		group.POST("/article/create", article.Create)
		group.POST("/article/list", article.List)
		group.GET("/article/detail", article.Detail)
		group.POST("/article/send_message", article.SendMessage)
	}
}

func SetSocketRouter(router *gin.Engine) {
	// Simple group: v2
	group := router.Group("/ws")
	{
		group.GET("/connect", socket.Connect)
	}
}
