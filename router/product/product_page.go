package product

import (
	"gomall/dal"
	"gomall/dal/mysql"
	"gomall/handler/product"
	"gomall/service"

	"github.com/gin-gonic/gin"
)

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *gin.Engine) {
	productDAL := dal.NewProductDAL(mysql.DB)
	product := product.NewProductHandler(service.NewProductService(productDAL))
	root := r.Group("/", RootMw()...)
	root.POST("/product", append(GetproductMw(), product.CreateProduct)...)
	root.GET("/search", append(SearchProductsMw(), product.SearchProducts)...)
	root.GET("/products", append(GetproductMw(), product.ListProducts)...)

}
