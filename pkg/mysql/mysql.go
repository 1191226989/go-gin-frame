package mysql

import (
	"fmt"
	"go-gin-frame/config"
	"sync"
	"time"

	"gorm.io/driver/mysql"
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
		cfg := config.Get().MySQL
		// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Pass, cfg.Addr, cfg.Name)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		sqlDB, err := db.DB()
		if err != nil {
			return
		}

		// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

		// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)

		// 设置最大连接超时
		sqlDB.SetConnMaxLifetime(time.Minute * cfg.ConnMaxLifeTime)
	})
	return db, err
}
