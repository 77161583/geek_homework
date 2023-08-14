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
	db := initDB()
	server := initWebServer()
	//注册路由
	u := initUser(db)
	u.RegisterRoutes(server)
	//启动
	server.Run(":8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	//跨域可以在这里处理...

	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		//只会在初始化的过程中panic
		//panic相当于整个goroutine结束
		//一旦初始化出错，应用就不要再启动了
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
