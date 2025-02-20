package cart

import (
	"gomall/handler/cart"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	// 购物车相关接口

	root := r.Group("/cart", RootMw()...)
	root.POST("/add", append(AddCartItemMw(), cart.AddtoCartItem)...)
	root.GET("/get", append(GetCartMw(), cart.GetCart)...)
	root.DELETE("/clear", cart.ClearCartHandler) // 清空购物车
	// cartGroup := r.Group("/cart")
	// {
	// 	cartGroup.POST("/add", handler.AddToCartHandler)     // 添加商品到购物车
	// 	cartGroup.GET("/get", handler.GetCartHandler)        // 获取购物车信息
	// 	cartGroup.DELETE("/clear", handler.ClearCartHandler) // 清空购物车
	// }
}
