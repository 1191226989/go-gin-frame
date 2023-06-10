package article

import "github.com/gin-gonic/gin"

func List(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List",
	})
}
