package userservice

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateToken(email string) (string, bool) {
	defer func() {
		fmt.Println("jwt token 创建异常", recover())
	}()
	ttl := time.Now().Add(time.Minute * 5)
	clamis := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ttl.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis)
	str, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	} else {
		return str, true
	}
}

func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	currentClaims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, currentClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, currentClaims, err
}
