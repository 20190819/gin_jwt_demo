package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func Migrations(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		logrus.Errorf("迁移文件报错：%v",err)
	}

}
