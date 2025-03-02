package order

import (
	"github.com/gin-gonic/gin"
	"gomall/handler/order"
)

func Register(r *gin.Engine) {

	//订单相关接口
	root := r.Group("/order", RootMw()...)
	{
		root.POST("/create", append(PlaceOrderMw(), order.PlaceOrderHandler)...)   //创建订单
		root.POST("/update", append(UpdateOrderMw(), order.UpdateOrderHandler)...) //修改订单信息
		//	root.DELETE("/cancel", append(StartOrderCancellationJobMw(), order.StartOrderCancellationJobHandler)...) //定时订单取消
	}

	//orderGroup := r.Group("/order")
	//{
	//	orderGroup.POST("/create", order.PlaceOrderHandler)  //创建订单
	//	orderGroup.POST("/update", order.UpdateOrderHandler) //修改订单信息
	//	//orderGroup.DELETE("/cancel", order.StartOrderCancellationJobHandler) //定时订单取消
	//
	//}
}
