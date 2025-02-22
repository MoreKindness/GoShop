package about

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {
	root := r.Group("/", RootMw()...)
	root.POST("/about", append(AboutMw())...)
}
