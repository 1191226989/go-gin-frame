package article

import (
	"go-gin-frame/internal/database/sqlite3"
	"sync"

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

var (
	m    *model
	once sync.Once
)

type model struct {
	db *gorm.DB
}

func NewModel() (*model, error) {
	var err error
	once.Do(func() {
		database, dbErr := sqlite3.NewDB()
		m = &model{
			db: database,
		}
		err = dbErr
	})
	return m, err
}
