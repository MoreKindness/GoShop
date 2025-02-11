package cart

import (
	"github.com/gin-gonic/gin"
	"gomall/handler/cart"
)

func Register(r *gin.Engine) {

	root := r.Group("/", RootMw()...)
	root.POST("/cart", append(AddCartItemMw(), cart.AddCartItem)...)
	root.GET("/cart", append(GetCartMw(), cart.GetCart)...)
}
