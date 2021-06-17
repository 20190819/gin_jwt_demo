package userservice

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var jwtKey = []byte("yangliang4488")

type Claims struct {
	Email string
	jwt.StandardClaims
}

func CreateToken(email string) string {
	ttl := time.Now().Add(time.Minute * 5)
	clamis := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ttl.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, clamis)
	str, err := token.SignedString(jwtKey)
	if err != nil {
		logrus.Errorf("token 创建失败:%v\n", err)
		return ""
	} else {
		return str
	}
}
