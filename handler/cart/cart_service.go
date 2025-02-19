package cart

import (
	"gomall/model"
	"gomall/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddToCartHandler 添加商品到购物车
func AddtoCartItem(c *gin.Context) {
	var item model.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数解析错误"})
		return
	}

	userID, _ := strconv.Atoi(c.Query("user_id"))

	if err := service.AddToCart(userID, item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加购物车失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "商品已添加到购物车"})
}

// GetCartHandler 获取购物车信息
func GetCart(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	cart, err := service.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "购物车为空"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

// ClearCartHandler 清空购物车
func ClearCartHandler(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	if err := service.ClearCart(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "清空购物车失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "购物车已清空"})
}
