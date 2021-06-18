package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	modelUser "github.com/yangliang4488/gin_jwt_demo/app/models/user"
	"github.com/yangliang4488/gin_jwt_demo/config/database"
	"net/http"
)

func Home(context *gin.Context) {
	claims,ok:=context.Get("claims")
	if !ok{
		logrus.Error("context 取 claims 失败")
		return
	}
	logrus.Info("claims",claims)

	mysqldb := database.MysqlDB
	users := []modelUser.User{}
	mysqldb.Find(&users)
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "操作成功",
		"data": users,
	})
}
