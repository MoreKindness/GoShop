package auth

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gomall/dal"
	"gomall/dal/dao"
	"gomall/dal/mysql"
	"gomall/model"
	"gomall/service"
	"net/http"
)

const (
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            *service.UserService
}

func NewUserHandler() *UserHandler {
	svc := service.NewUserService(dal.NewUserRepository(dao.NewUserDAO(mysql.DB)))
	return &UserHandler{
		emailRexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:            svc,
	}
}

// 注册路由组
//func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
//	//ug := server.Group("/auth")
//	//ug.GET("/register", h.SignUpPage) // 显示注册页面
//	//ug.POST("/register", h.SignUp)    // 处理注册表单
//	//ug.GET("/login", h.LoginPage)     // 显示注册页面
//	//ug.POST("/login", h.Login)        // 处理注册表单
//	//ug.GET("/home", h.HomePage)       // 主界面路由
//}

// Register .
// @router /auth/register [POST]

func (h *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	// 表单绑定
	var req SignUpReq
	req.Email = ctx.PostForm("email")
	req.Password = ctx.PostForm("password")
	req.ConfirmPassword = ctx.PostForm("password-confirm")
	pageName := "sign-up"

	// 校验邮箱格式
	isEmail, err := h.emailRexExp.MatchString(req.Email)
	if err != nil {
		RenderPageMsg(ctx, pageName, "系统错误")
		return
	}
	if !isEmail {
		RenderPageMsg(ctx, pageName, "非法邮箱格式")
		return
	}

	// 校验密码一致性
	if req.Password != req.ConfirmPassword {
		RenderPageMsg(ctx, pageName, "两次输入密码不一致")
		return
	}

	// 校验密码强度
	isPassword, err := h.passwordRexExp.MatchString(req.Password)
	if err != nil {
		RenderPageMsg(ctx, pageName, "系统错误")
		return
	}
	if !isPassword {
		RenderPageMsg(ctx, pageName, "密码必须包含字母、数字、特殊字符，并且不少于八位")
		return
	}

	// 调用注册服务
	err = h.svc.Signup(ctx, model.User{
		Email:    req.Email,
		Password: req.Password,
	})
	switch err {
	case nil:
		RenderPageMsg(ctx, pageName, "注册成功！")
	case service.ErrDuplicateEmail:
		RenderPageMsg(ctx, pageName, "邮箱已被注册，请换一个")
	default:
		RenderPageMsg(ctx, pageName, "系统错误")
	}
}

// 渲染页面信息
func RenderPageMsg(ctx *gin.Context, templateName string, msg string) {
	ctx.HTML(http.StatusOK, templateName, gin.H{
		"error": msg,
	})
}

// Login .
// @router /auth/login [POST]
func (h *UserHandler) Login(ctx *gin.Context) {
	type Req struct {
		Email    string
		Password string
	}

	var req Req

	req.Email = ctx.PostForm("email")
	req.Password = ctx.PostForm("password")
	pageName := "sign-in"
	u, err := h.svc.Login(ctx, req.Email, req.Password)
	switch err {
	case nil:
		sess := sessions.Default(ctx)
		sess.Set("user_id", u.ID)
		user_cat, err := service.GetCart(u.ID)
		if err != nil {
			RenderPageMsg(ctx, pageName, "系统错误！")
		}
		sess.Set("cart", user_cat)
		//sess.Options(sessions.Options{MaxAge: 900})
		err = sess.Save()
		if err != nil {
			RenderPageMsg(ctx, pageName, "系统错误！")
			return
		}
		ctx.Redirect(http.StatusFound, "/") // 登录成功后重定向到主页
		return
	case service.ErrInvalidUserOrPassword:
		RenderPageMsg(ctx, pageName, "用户名或密码错误！")
	default:
		RenderPageMsg(ctx, pageName, "用户未注册")

	}
}

// Logout .
// @router /auth/logout [POST]
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user_id")
	session.Delete("cart")
	session.Save()
	c.Redirect(http.StatusFound, "/")
}
