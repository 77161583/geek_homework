package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"home2/internal/repository"
	"home2/internal/repository/dao"
	"home2/internal/service"
	"home2/internal/web"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		//只会在初始化的过程中panic
		//panic相当于整个goroutine结束
		//一旦初始化出错，应用就不要再启动了
		panic(err)
	}
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	//加载user模块
	server := gin.Default()
	u := web.NewUserHandler(svc)
	u.RegisterRoutes(server)
	server.Run(":8080")
}
