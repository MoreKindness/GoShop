package product

import (
	"gomall/handler/product"
	"gomall/service"

	"github.com/gin-gonic/gin"
)

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *gin.Engine) {
	product := product.NewProductHandler(service.NewProductService())
	root := r.Group("/", RootMw()...)
	root.POST("/product", append(GetproductMw(), product.CreateProduct)...)
	root.GET("/search", append(SearchProductsMw(), product.SearchProducts)...)
	root.GET("/products", append(GetproductMw(), product.ListProducts)...)
	root.GET("/product", append(GetproductMw(), product.SearchProducts)...)
}
