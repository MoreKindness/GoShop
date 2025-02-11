package product

import (
	"github.com/gin-gonic/gin"
	"gomall/handler/product"
)

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *gin.Engine) {

	root := r.Group("/", RootMw()...)
	root.GET("/product", append(GetproductMw(), product.GetProduct)...)
	root.GET("/search", append(SearchproducsMw(), product.SearchProducs)...)
}
