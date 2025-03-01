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

	db, err := mysql.Init()
	if err != nil {
		panic(err)
	}

	productDAL := dal.NewProductDAL(db)
	product := product.NewProductHandler(service.NewProductService(productDAL))
	root := r.Group("/", RootMw()...)
	root.POST("/product", append(GetproductMw(), product.CreateProduct)...)
	root.GET("/search", append(SearchProductsMw(), product.SearchProducts)...)
	root.GET("/products", append(GetproductMw(), product.ListProducts)...)

}
