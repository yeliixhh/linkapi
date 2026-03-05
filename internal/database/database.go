package database

import (
	"fmt"
	"time"

	"github.com/yeliixhh/linkapi/internal/config"
	"github.com/yeliixhh/linkapi/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 初始化DB
func NewDB(config *config.Config) (*gorm.DB, error) {
	logger.Log.Info(fmt.Sprintf("连接数据库 host: %s", "localhost"))

	// https://github.com/go-gorm/postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBConf.Host,
		config.DBConf.User,
		config.DBConf.Password,
		config.DBConf.DBName,
		config.DBConf.Port,
	)

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
