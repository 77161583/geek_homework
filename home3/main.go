package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"home3/internal/repository"
	"home3/internal/repository/dao"
	"home3/internal/service"
	"home3/internal/web"
	middleware "home3/internal/web/middleware"
	"net/http"
	"strings"
	"time"
)

func main() {
	db := initDB()
	server := initWebServer()
	//注册路由
	u := initUser(db)
	u.RegisterRoutes(server)
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "来了老弟！")
	})
	//启动
	server.Run(":8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	//跨域可以在这里处理...
	server.Use(func(ctx *gin.Context) {
		fmt.Println("The first middleware")
	})

	server.Use(func(ctx *gin.Context) {
		fmt.Println("The second middleware")
	})

	//redisClient := redis.NewClient(&redis.Options{
	//	Addr: config.Config.Redis.Addr,
	//})
	//server.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())

	server.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,                    // 是否允许你带 cookie 之类的东西
		ExposeHeaders:    []string{"x-jwt-token"}, //不设置这个，前端读不到
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				//你的开发环境
				return true
			}
			return strings.Contains(origin, "your company.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	//登陆之后保存登陆信息 步骤1
	//store := cookie.NewStore([]byte("secret"))
	//store := memstore.NewStore([]byte("WbeWraNhhon7NxWP7w9WSKMLzZ8cTiwM"), []byte("YHgJ7VQuszth64EuHphVSYVN9SY9NA76"))
	//第一个参数是最大连接数量，面试如果问具体是多少，回答“压力测试，性能测试”
	//第二个参数连接方式
	//第三第四，端口号，密码
	//五六就是两个key
	store := memstore.NewStore([]byte("WbeWraNhhon7NxWP7w9WSKMLzZ8cTiwM"),
		[]byte("YHgJ7VQuszth64EuHphVSYVN9SY9NA76"))
	//store, err := redis.NewStore(16, "tcp", "localhost:6379", "",
	//	[]byte("WbeWraNhhon7NxWP7w9WSKMLzZ8cTiwM"), []byte("YHgJ7VQuszth64EuHphVSYVN9SY9NA76"))
	//if err != nil {
	//	panic(err)
	//}

	//myStore := &sqlx_store.Store{}
	server.Use(sessions.Sessions("mysession", store))

	//登陆之后的校验 - 登陆之后保存登陆信息 步骤3
	//server.Use(middleware.NewLoginMiddlewareBuilder().
	//	IgnorePaths("/users/login").
	//	IgnorePaths("/users/signup").Build())

	server.Use(middleware.NewLoginJWTMiddlewareBuilder().
		IgnorePaths("/users/login").
		IgnorePaths("/users/signup").Build())
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
