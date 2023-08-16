package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"home2/internal/domain"
	"home2/internal/service"
	"net/http"
)

// UserHandler 定义和跟用户有关的路由
type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	birthdayExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	//校验参数
	const (
		emailRegExpPattern    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
		passwordRegExpPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
		birthdayRegExpPattern = `^\d{4}-\d{2}-\d{2}$`
	)
	emailRegExp := regexp.MustCompile(emailRegExpPattern, 0)
	passwordRegExp := regexp.MustCompile(passwordRegExpPattern, 0)
	birthdayRegExp := regexp.MustCompile(birthdayRegExpPattern, 0)
	return &UserHandler{
		svc:         svc,
		emailExp:    emailRegExp,
		passwordExp: passwordRegExp,
		birthdayExp: birthdayRegExp,
	}
}

func (u *UserHandler) RegisterRoutes(serve *gin.Engine) {
	ug := serve.Group("/users")
	ug.POST("login", u.Login)
	ug.POST("signup", u.SignUp)
	ug.POST("edit", u.Edit)
	ug.POST("profile", u.Profile)
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

	//调用service 层了
	err = u.svc.SignUp(ctx, domain.User{Email: req.Email, Password: req.Password})
	if err == service.ErrUseDuplicateEmail {
		ctx.String(http.StatusOK, "邮箱冲突")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
}

func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "用户名/邮箱或密码不对")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	//登陆成功之后获取sessions
	//登陆之后保存登陆信息 步骤2
	sess := sessions.Default(ctx)
	//设置seesion的值
	sess.Set("userId", user.Id)
	sess.Save()
	ctx.String(http.StatusOK, "登陆成功")
	return

}

// Edit 作业 根据用户id修改数据
func (u *UserHandler) Edit(ctx *gin.Context) {
	//定义接收数据
	type UserDataReq struct {
		UserId       int64  `json:"userId"`
		NickName     string `json:"nickName"`
		Birthday     string `json:"birthday"`
		Introduction string `json:"introduction"`
	}
	//实例化一个req
	var req UserDataReq
	//用bing 获取数据
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求数据绑定失败"})
	}
	if req.UserId == 0 {
		ctx.String(http.StatusOK, "用户Id丢失，无法编辑")
		return
	}
	//查看该条用户是否存在
	UserData, _ := u.svc.FindById(ctx, req.UserId)
	if UserData.Id == 0 {
		ctx.String(http.StatusOK, "没有改用户，无法编辑")
		return
	}
	//参数验证
	if req.NickName != "" {
		if len(req.NickName) < 6 || len(req.NickName) > 30 {
			ctx.String(http.StatusOK, "昵称大小需要保持2~10个汉字")
			return
		}
	}
	if req.Birthday != "" {
		isBirthday, _ := u.birthdayExp.MatchString(req.Birthday)
		if !isBirthday {
			ctx.String(http.StatusOK, "生日日期格式不正确，应为YYYY-MM-DD格式")
			return
		}
		if req.Introduction == "" {
			ctx.String(http.StatusOK, "个人简介不能为空")
			return
		}
	}
	//调用service
	err := u.svc.Edit(ctx, domain.User{
		Id:           req.UserId,
		NickName:     req.NickName,
		Birthday:     req.Birthday,
		Introduction: req.Introduction,
	})
	if err != nil {
		ctx.String(http.StatusOK, "修改失败")
		return
	}
	//ctx.JSON(http.StatusOK, UserData)
	ctx.String(http.StatusOK, "修改成功")
}

// Profile 作业回显
func (u *UserHandler) Profile(ctx *gin.Context) {
	//定义接收数据
	type UserAllDataReq struct {
		UserId int64 `json:"userId"`
	}
	//实例化一个req
	var reqSelect UserAllDataReq
	//用bing 获取数据
	if err := ctx.Bind(&reqSelect); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求数据绑定失败"})
	}
	if reqSelect.UserId == 0 {
		ctx.String(http.StatusOK, "用户Id丢失，无法编辑")
		return
	}
	//查看该条用户是否存在
	UserData, _ := u.svc.FindById(ctx, reqSelect.UserId)
	if UserData.Id == 0 {
		ctx.String(http.StatusOK, "未找到该用户")
		return
	}
	ctx.JSON(http.StatusOK, UserData)
}
