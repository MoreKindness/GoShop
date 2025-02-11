package auth

import (
	"github.com/gin-gonic/gin"
	"gomall/handler/auth"
)

func Register(r *gin.Engine) {

	root := r.Group("/", RootMw()...)
	{
		_auth := root.Group("/auth", AuthMw()...)
		_auth.POST("/login", append(LoginMw(), auth.Login)...)
		_auth.POST("/logout", append(LogoutMw(), auth.Logout)...)
		_auth.POST("/register", append(RegisterMw(), auth.Register)...)
	}
}
