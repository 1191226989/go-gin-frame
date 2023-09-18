package article

import (
	"go-gin-frame/pkg/sqlite3"
	"sync"

	"gorm.io/gorm"
)

var (
	err  error
	m    *model
	once sync.Once
)

type model struct {
	db *gorm.DB
}

func NewModel() (*model, error) {
	once.Do(func() {
		database, dbErr := sqlite3.GetInstance()
		m = &model{
			db: database,
		}
		err = dbErr
	})
	return m, err
}

// 文章详情
func (m *model) Detail(id int) (*Article, error) {
	article := Article{}
	tx := m.db.First(&article, id)
	if tx.Error != nil {
		return &article, tx.Error
	}
	return &article, nil
}
