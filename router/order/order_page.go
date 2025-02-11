package order

import (
	"github.com/gin-gonic/gin"
	"gomall/handler/order"
)

func Register(r *gin.Engine) {

	root := r.Group("/", RootMw()...)
	root.GET("/order", append(OrderlistMw(), order.OrderList)...)
}
