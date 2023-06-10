package article

import "github.com/gin-gonic/gin"

func Detail(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Detail",
	})
}
