package cart

import (
	"gomall/model"
	"gomall/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获取购物车
func GetCartHandler(c *gin.Context) {
	cartID, err := strconv.ParseUint(c.Query("cart_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart_id"})
		return
	}

	cart, err := service.GetCart(uint(cartID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

//添加商品到购物车

// AddToCartHandler 处理添加商品到购物车的请求
func AddToCartHandler(c *gin.Context) {
	var request struct {
		CartID    uint `json:"cart_id"`
		ProductID uint `json:"product_id"`
		Quantity  uint `json:"quantity"`
	}

	// 解析 JSON 数据
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	// 检查 `cart_id` 是否存在
	_, err := service.GetCart(request.CartID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
		}
		return
	}

	// 创建 `CartItem` 对象
	item := model.CartItem{
		CartID:    request.CartID,
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
	}

	// 调用 `service.AddToCart`
	err = service.AddToCart(request.CartID, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added successfully"})
}

// 空购物车
func ClearCartHandler(c *gin.Context) {
	var request struct {
		CartID uint `json:"cart_id" form:"cart_id"`
	}

	// 绑定 JSON 或表单数据
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart_id"})
		return
	}

	err := service.ClearCart(request.CartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared"})
}
