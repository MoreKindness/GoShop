package router

import (
	"github.com/gin-gonic/gin"
	"gomall/router/about"
	"gomall/router/auth"
	"gomall/router/cart"
	"gomall/router/category"
	"gomall/router/checkout"
	"gomall/router/home"
	"gomall/router/order"
	"gomall/router/product"
)

func GeneratedRegister(r *gin.Engine) {
	about.Register(r)

	order.Register(r)

	checkout.Register(r)

	auth.Register(r)

	cart.Register(r)

	category.Register(r)

	product.Register(r)

	home.Register(r)
}
