package category

import (
	"github.com/gin-gonic/gin"
	"gomall/handler/category"
)

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *gin.Engine) {

	root := r.Group("/", RootMw()...)
	{
		_category := root.Group("/category", CategoryMw()...)
		_category.GET("/:category", append(Category0Mw(), category.Category)...)
	}
}
