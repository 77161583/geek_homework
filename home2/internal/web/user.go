package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler 定义和跟用户有关的路由
type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	//校验参数
	const (
		emailRegExpPattern    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
		passwordRegExpPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailRegExp := regexp.MustCompile(emailRegExpPattern, 0)
	passwordRegExp := regexp.MustCompile(passwordRegExpPattern, 0)
	return &UserHandler{
		emailExp:    emailRegExp,
		passwordExp: passwordRegExp,
	}
}

func (u *UserHandler) RegisterRoutes(serve *gin.Engine) {
	ug := serve.Group("/users")
	ug.POST("login", u.Login)
	ug.POST("signup", u.SignUp)
	ug.POST("edit", u.Edit)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	var req SignUpReq
	//Bind 方法回根据 Content-type 来解析你的数据到req 里面
	//解析错了会直接返回400 的错误
	if err := ctx.Bind(&req); err != nil {
		return
	}

	isEmail, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "Email格式不正确")
		return
	}
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次密码输入不一样！")
		return
	}
	isPassword, err := u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误1")
		return
	}
	if !isPassword {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字，特殊字符")
		return
	}
	ctx.String(http.StatusOK, "注册成功")
}

func (u *UserHandler) Login(ctx *gin.Context) {

}

func (u *UserHandler) Edit(ctx *gin.Context) {

}
