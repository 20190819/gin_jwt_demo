package main

import (
	"fmt"

	"github.com/yangliang4488/gin_jwt_demo/config"
	_ "github.com/yangliang4488/gin_jwt_demo/config"
)

func main() {
	fmt.Printf("mysql db 类型是:%T\n", config.MysqlDB)
}
