package main

import (
	"fmt"
	"github.com/yangliang4488/gin_jwt_demo/config/database"
	_ "github.com/yangliang4488/gin_jwt_demo/routes"
)

func main() {
	fmt.Printf("mysql db 类型是:%T\n", database.MysqlDB)
}
