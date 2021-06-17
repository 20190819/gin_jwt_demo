package user

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"column:email;type:varchar(255);not null;unique;"`
	Username string
	Password string
}

func Hash(password string) string {
	byteStr, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logrus.Errorf("密码加密出错:%v\n", err)
	}
	return string(byteStr)
}

func Checkhash(password string, hashPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(password))
	return err == nil
}
