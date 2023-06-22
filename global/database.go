package global

import (
	"go-gin-frame/internal/database/sqlite3"
	"go-gin-frame/internal/model/article"
	"sync"
)

// 组件初始化

var once sync.Once

func Init() {
	once.Do(func() {
		initSqlite()
		initMysql()
		initRedis()
	})
}

func initSqlite() {
	db, err := sqlite3.NewDB()
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&article.Article{})
}

func initMysql() {

}

func initRedis() {

}
