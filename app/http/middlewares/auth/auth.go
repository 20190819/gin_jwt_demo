package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	serviceUser "github.com/yangliang4488/gin_jwt_demo/app/services/user_service"
	"net/http"
)

func Api() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.Request.Header.Get("token")
		if tokenStr == "" {
			context.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  "未携带指定 header 参数 token",
			})
			return
		}
		logrus.Info("has token")
		// 解析 token
		_, claims, err := serviceUser.ParseToken(tokenStr)
		if err != nil {
			context.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  "登录检验失败~",
				"err":  err.Error(),
			})
			return
		} else {
			context.Set("claims", claims)
		}
	}
}
