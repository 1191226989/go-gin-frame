package initial

import (
	"go-gin-frame/internal/model/article"
	"go-gin-frame/pkg/mysql"
	"go-gin-frame/pkg/redis"
	"go-gin-frame/pkg/sqlite3"

	"github.com/sirupsen/logrus"
)

func InitSqlite() {
	db, err := sqlite3.GetInstance()
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&article.Article{})
}

func InitMysql() {
	db, err := mysql.GetInstance()
	if err != nil {
		panic(err)
	}
	logrus.Infoln(db)
}

func InitRedis() {
	c, err := redis.GetInstance()
	if err != nil {
		panic(err)
	}
	logrus.Infoln(c)
}
