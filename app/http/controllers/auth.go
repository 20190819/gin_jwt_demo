package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	modelUser "github.com/yangliang4488/gin_jwt_demo/app/models/user"
	serviceUser "github.com/yangliang4488/gin_jwt_demo/app/services/user_service"
	"github.com/yangliang4488/gin_jwt_demo/config"
)

// 注册
func Register(ctx *gin.Context) {
	user := modelUser.User{}
	mysqldb := config.MysqlDB

	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	if email == "" || password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "信息不完整",
		})
		return
	} else {
		mysqldb.Where("email=?", email).First(&user)

		if user.Email != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "邮箱已被占用",
			})
			return
		}
		// 注册
		res := mysqldb.Create(&user)
		if res.RowsAffected > 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "注册成功",
			})
			return
		}
	}
}

// 登录
func Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	if email == "" || password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "信息不完整",
		})
		return
	} else {
		user := modelUser.User{}
		mysqldb := config.MysqlDB
		mysqldb.Where("email=?", email).First(&user)
		res := modelUser.Checkhash(password, user.Password)
		if !res || user.Email == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		} else {
			// 登录成功，生成 token
			token := serviceUser.CreateToken(email)
			// 返回结果
			ctx.JSON(http.StatusOK, gin.H{
				"code":  -1,
				"msg":   "信息不完整",
				"token": token,
				"user":  user,
			})
		}

	}
}

// 生成 token
func CreateToken(ctx *gin.Context) {
	//
}
