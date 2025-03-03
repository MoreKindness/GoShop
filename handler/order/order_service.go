package order

import (
	"gomall/model"
	"gomall/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Request 返回tmpl的参数结构体
type Request struct {
	CreatedDate string  `json:"created_date"`
	OrderId     uint    `json:"id"`
	Picture     string  `json:"picture"`
	ProductName string  `json:"product_name"`
	Qty         int32   `json:"quantity"`
	Cost        float64 `json:"price"`
}

// OrdersResponse 包含多个订单的响应
type OrdersResponse struct {
	Orders []Request `json:"orders"`
}

// Response 获取订单返回数据
func Response(orders []model.Order) OrdersResponse {
	var resp OrdersResponse
	for _, order := range orders {
		req := Request{
			CreatedDate: order.CreatedAt.Format("2006-01-02"),
			OrderId:     order.ID,
			Picture:     order.OrderItems[0].Picture,
			ProductName: order.OrderItems[0].ProductName,
			Qty:         order.OrderItems[0].Quantity,
			Cost:        order.OrderItems[0].Price,
		}
		resp.Orders = append(resp.Orders, req)
	}
	return resp
}

// ResponseOne 获取订单返回数据
func ResponseOne(order model.Order) Request {
	var resp Request
	resp.CreatedDate = order.CreatedAt.Format("2006-01-02")
	resp.OrderId = order.ID
	resp.Picture = order.OrderItems[0].Picture
	resp.ProductName = order.OrderItems[0].ProductName
	resp.Qty = order.OrderItems[0].Quantity
	resp.Cost = order.OrderItems[0].Price
	return resp
}

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

	//JSON返回订单信息
	//userID := c.GetInt("user_id")
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Order created successfully",
	//	"user_id": userID,
	//})

	// HTML 将订单数据传递给模板
	c.HTML(http.StatusOK, "order", ResponseOne(order))
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
	updatedAt := time.Now()
	err := service.UpdateOrder(req.ID, uint32(userID), &req.Consignee, updatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order, err := service.ShowOrder(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// HTML 将订单数据传递给模板
	c.HTML(http.StatusOK, "order", ResponseOne(*order))

	// JSON 获取订单数据
	//c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

// ListOrdersHandler 处理获取所有订单的请求
func ListOrdersHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	orders, err := service.GetAllOrders(uint32(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// HTML 将订单数据传递给模板
	c.HTML(http.StatusOK, "order", Response(orders))
}

// CancelOrderHandler 处理取消订单的请求
func CancelOrderHandler(c *gin.Context) {
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	err = service.CancelOrder(uint(ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order, err := service.ShowOrder(uint(ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// HTML 将订单数据传递给模板
	c.HTML(http.StatusOK, "order", ResponseOne(*order))

	// JSON 获取订单数据
	//c.JSON(http.StatusOK, gin.H{"message": "Order canceled successfully"})
}

//// StartOrderCancellationJobHandler 定时取消订单
//func StartOrderCancellationJobHandler(c *gin.Context) {
//	go service.CancelExpiredOrders()
//	c.JSON(http.StatusOK, gin.H{"message": "The scheduled order was cancelled successfully"})
//}
