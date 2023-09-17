package sqlite3

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	err  error
	db   *gorm.DB
	once sync.Once
)

// 单例
func GetInstance() (*gorm.DB, error) {
	once.Do(func() {
		db, err = gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
	})
	return db, err
}
