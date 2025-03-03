package auth

import (
	"github.com/gin-gonic/gin"
	"gomall/handler/auth"
)

func Register(r *gin.Engine) {
	// auth
	authH := auth.NewUserHandler()
	root := r.Group("/", RootMw()...)
	{
		_auth := root.Group("/auth", AuthMw()...)
		_auth.POST("/login", append(LoginMw(), authH.Login)...)
		_auth.POST("/logout", append(LogoutMw(), auth.Logout)...)
		_auth.POST("/register", append(RegisterMw(), authH.SignUp)...)
	}
}
