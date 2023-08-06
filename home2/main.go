package main

import (
	"github.com/gin-gonic/gin"
	"home2/internal/web"
)

func main() {
	server := gin.Default()
	//加载user模块
	u := web.NewUserHandler()
	u.RegisterRoutes(server)

	server.Run(":8080")
}
