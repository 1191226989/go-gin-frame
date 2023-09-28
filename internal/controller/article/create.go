package article

import (
	"go-gin-frame/internal/code"
	"go-gin-frame/internal/service/article"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Create(c *gin.Context) {
	data := &article.CreateArticleData{
		Title:   "test",
		Content: "test",
	}
	s := article.NewService()
	id, err := s.Create(data)
	if err != nil {
		c.JSON(code.ArticleCreateError, gin.H{
			"message": err,
		})
	}

	logrus.Infoln("article id: ", id)
	c.JSON(200, gin.H{
		"message": "Create Success",
	})
}
