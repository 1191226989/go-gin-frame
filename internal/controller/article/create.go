package article

import "github.com/gin-gonic/gin"

func Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create",
	})
}
