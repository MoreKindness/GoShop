package home

import (
	"github.com/gin-gonic/gin"
	"gomall/handler/home"
)

func Register(r *gin.Engine) {

	root := r.Group("/", RootMw()...)
	root.GET("/", append(HomeMw(), home.Home)...)
}
