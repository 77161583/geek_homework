package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
}

func (u *UserHandler) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "请求成功")
}
func (u *UserHandler) Login(ctx *gin.Context) {

}
func (u *UserHandler) Signup(ctx *gin.Context) {

}
func (u *UserHandler) GetById(ctx *gin.Context) {
	userId := ctx.Param("id")
	ctx.String(http.StatusOK, "用户Id是："+userId)
}
func SetupUserRoutes(r *gin.Engine, userUrl urlHandler) {
	user := r.Group("/user")
	{
		user.GET("/", userUrl.Index)
		user.POST("/", userUrl.Login)
		user.POST("/", userUrl.Signup)
		user.GET("/:id", userUrl.GetById)
	}
}
