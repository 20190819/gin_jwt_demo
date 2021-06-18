package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yangliang4488/gin_jwt_demo/app/http/controllers"
	middlewareAuth "github.com/yangliang4488/gin_jwt_demo/app/http/middlewares/auth"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 注册中间件
	r.Use(middlewareAuth.Api())

	r.POST("/register",controllers.Register)
	r.POST("/login",controllers.Login)
	return r
}
func init(){
	r := SetupRouter()
	if err := r.Run(":9090"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
