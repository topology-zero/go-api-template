package model

import (
	"fmt"
	"time"

	"go-api-template/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Setup() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlConf.User,
		config.MysqlConf.Pwd,
		config.MysqlConf.Host,
		config.MysqlConf.Port,
		config.MysqlConf.Db,
	)
	var err error

	logMode := glog.Silent
	if config.Env == "local" || config.Env == "dev" {
		logMode = glog.Info
	}
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: NewLogger(glog.Config{
			SlowThreshold:             200 * time.Millisecond,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logMode,
		}),
	})
	if err != nil {
		logrus.Panic(err)
	}
}

// DB 获取数据库实例
func DB() *gorm.DB {
	return db
}
