package main

import (
	"github.com/joho/godotenv"
	"github.com/polarrana/go_task4/config"
	"github.com/polarrana/go_task4/routes"
	"log"
)

func main() {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or error loading .env file")
	}
	// 初始化数据库
	db := config.InitDB()

	// 设置路由
	r := routes.SetupRoutes(db)

	// 启动服务器
	log.Println("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
