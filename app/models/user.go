package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"column:email;type:varchar(255);not null;unique;"`
	Username string
	Password string
}
