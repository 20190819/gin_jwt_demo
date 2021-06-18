package database

import (
	"time"

	"github.com/sirupsen/logrus"
	modelUser "github.com/yangliang4488/gin_jwt_demo/app/models/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

//var err error
func init() {
	// 连接
	mysqlDB := connections()
	MysqlDB = mysqlDB
	// 迁移
	migrations(mysqlDB)
}

func connections() (mysqlDB *gorm.DB) {
	config := mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3308)/gin_jwt_demo?charset=utf8&parseTime=True&loc=Local",
	})
	mysqlDB, err := gorm.Open(config, &gorm.Config{})
	if err != nil {
		logrus.Error("连接 mysql 出错")
		return
	}
	sqlDB, _ := mysqlDB.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	return
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&modelUser.User{})
	if err != nil {
		logrus.Error("migrations 出错")
		return
	}
}
