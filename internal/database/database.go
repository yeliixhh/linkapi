package database

import (
	"fmt"
	"time"

	"github.com/yeliixhh/linkapi/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 初始化DB
func NewDB() (*gorm.DB, error) {
	logger.Log.Info(fmt.Sprintf("连接数据库 host: %s", "localhost"))

	// https://github.com/go-gorm/postgres
	dsn := "host=localhost user=root password=123456 dbname=link_api port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	s, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量
	s.SetMaxOpenConns(10)
	s.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	s.SetConnMaxLifetime(time.Hour)

	return db, err
}
