package article

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	Title   string
	Content string
}

// 表名称
func (a *Article) Table() string {
	return "article"
}
