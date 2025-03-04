package order

import (
	"github.com/gin-gonic/gin"
	"gomall/handler/order"
)

func Register(r *gin.Engine) {

	//订单相关接口
	root := r.Group("/order", RootMw()...)
	{
		root.POST("/create", append(PlaceOrderMw(), order.PlaceOrderHandler)...)       //创建订单
		root.POST("/update", append(UpdateOrderMw(), order.UpdateOrderHandler)...)     //修改订单信息
		root.GET("/", append(ListOrdersMw(), order.ListOrdersHandler)...)              // 列出所有订单信息
		root.POST("/cancel/:id", append(CancelOrderMw(), order.CancelOrderHandler)...) // 取消订单
	}
}
