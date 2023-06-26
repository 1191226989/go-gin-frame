package article

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin-frame/assets"
)

func Download(c *gin.Context) {
	// 下载嵌入静态资源
	c.FileFromFS("static/avatar.png", http.FS(assets.Static))
}
