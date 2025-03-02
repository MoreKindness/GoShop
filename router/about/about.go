package about

import (
	"github.com/gin-gonic/gin"
	"gomall/handler/about"
)

func Register(r *gin.Engine) {
	root := r.Group("/", RootMw()...)
	root.POST("/about", append(AboutMw(), about.About)...)
}
