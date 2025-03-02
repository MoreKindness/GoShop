package order

import (
	"gomall/model"
	"gomall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//// OrderList .
//// @router /order [GET]
//func OrderList(c *gin.Context) {
//	// user_id, err := c.Cookie("user_id")
//	// if err != nil {
//	// 	c.String(http.StatusForbidden, err.Error())
//	// }
//	// cart_data, err := c.Cookie("cart_data")
//	// if err != nil {
//	// 	c.String(http.StatusForbidden, err.Error())
//	// }
//	c.HTML(http.StatusOK, "order", gin.H{})
//}

// PlaceOrderHandler 处理创建订单的HTTP请求
func PlaceOrderHandler(c *gin.Context) {
	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.PlaceOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

// UpdateOrderHandler 处理修改订单信息请求
func UpdateOrderHandler(c *gin.Context) {
	var req struct {
		ID        uint            `json:"id"`
		Consignee model.Consignee `json:"consignee"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt("user_id")
	err := service.UpdateOrder(req.ID, uint32(userID), &req.Consignee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

//// StartOrderCancellationJobHandler 定时取消订单
//func StartOrderCancellationJobHandler(c *gin.Context) {
//	go service.CancelExpiredOrders()
//	c.JSON(http.StatusOK, gin.H{"message": "The scheduled order was cancelled successfully"})
//}
