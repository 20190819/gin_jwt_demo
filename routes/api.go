package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yangliang4488/gin_jwt_demo/app/http/controllers"
	middlewareAuth "github.com/yangliang4488/gin_jwt_demo/app/http/middlewares/auth"
	"github.com/yangliang4488/gin_jwt_demo/config/app"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 注册中间件
	demo := r.Group("/", middlewareAuth.Api())
	{
		r.POST("/register", controllers.Register)
		r.POST("/login", controllers.Login)
		demo.GET("/home", controllers.Home)
	}

	return r
}
func init() {
	r := SetupRouter()
	if err := r.Run(":" + app.Config("app", "SERVER_PORT")); err != nil {
		logrus.Errorf("startup service failed, err:%v\n", err)
	}
}
