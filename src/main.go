package main

import (
	"news/conf"
	"news/router"

	"os"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := router.NewRouter()
	port := os.Getenv("PORT")
	r.Run(":" + port)
}