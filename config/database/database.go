package database

import (
	"github.com/sirupsen/logrus"
	modelUser "github.com/yangliang4488/gin_jwt_demo/app/models/user"
	"github.com/yangliang4488/gin_jwt_demo/config/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
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

var dbDriver string =app.Config("database","DRIVER")
var dbHost string =app.Config("database","HOST")
var dbPort string =app.Config("database","PORT")
var dbUsername string =app.Config("database","USERNAME")
var dbPassword string =app.Config("database","PASSWORD")
var dbDatabaseName string =app.Config("database","database")

func connections() (mysqlDB *gorm.DB) {
	config := mysql.New(mysql.Config{
		DriverName:dbDriver,
		DSN: dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbDatabaseName+"?charset=utf8&parseTime=True&loc=Local",
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
