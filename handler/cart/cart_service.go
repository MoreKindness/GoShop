package cart

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gomall/model"
	"gomall/service"
	"net/http"
)

// 获取购物车
func GetCartHandler(c *gin.Context) {
	report := make(map[string]interface{})
	session := sessions.Default(c)
	cart := session.Get("cart")
	var _cart model.Cart
	if cart != nil {
		_cart = cart.(model.Cart)
		report["items"] = _cart.Items
	} else {
		c.HTML(http.StatusBadRequest, "home", gin.H{"error": "Cart not found"})
	}
	user := session.Get("user_id")
	if user != nil {
		report["user_id"] = user
	}
	var total float64 = 0.0
	for i, _ := range _cart.Items {
		total += float64(_cart.Items[i].Quantity) * _cart.Items[i].Price
	}
	report["total"] = total
	c.HTML(http.StatusOK, "cart", report)
}

//添加商品到购物车

// AddToCartHandler 处理添加商品到购物车的请求
func AddToCartHandler(c *gin.Context) {
	report := make(map[string]interface{})
	type CartRequest struct {
		ProductID  uint `form:"productId" binding:"required"`
		ProductNum uint `form:"productNum" binding:"required"`
	}

	var request CartRequest
	// 解析数据
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}
	session := sessions.Default(c)
	cart := session.Get("cart")
	user := session.Get("user_id")
	if user != nil {
		report["user_id"] = user
	}
	var _cart model.Cart
	if cart != nil {
		_cart = cart.(model.Cart)

	}
	// 检查 `cart_id` 是否存在
	_, err := service.GetCart(_cart.ID)
	if err != nil {
		report["error"] = "Failed to get cart: " + err.Error()
		c.HTML(http.StatusBadRequest, "cart", report)
		return
	}

	product, err := service.NewProductService().GetProductByID(int(request.ProductID))
	if err != nil {
		report["error"] = "Failed to get product: " + err.Error()
		c.HTML(http.StatusBadRequest, "cart", report)
		return
	}
	// 创建 `CartItem` 对象
	item := model.CartItem{
		CartID:      _cart.ID,
		ProductID:   request.ProductID,
		Name:        product.Name,
		Picture:     product.Picture,
		Quantity:    request.ProductNum,
		Description: product.Description,
		Price:       product.Price,
	}

	// 调用 `service.AddToCart`
	item, err = service.AddToCart(_cart.ID, item)
	if err != nil {
		report["error"] = "Failed to add item to cart: " + err.Error()
		c.HTML(http.StatusBadRequest, "cart", report)
		return
	}
	_cart.Items = append(_cart.Items, item)
	session.Set("cart", _cart)
	report["cart_num"] = len(_cart.Items)
	var total float64 = 0.0
	for i, _ := range _cart.Items {
		total += float64(_cart.Items[i].Quantity) * _cart.Items[i].Price
	}
	report["total"] = total
	report["items"] = _cart.Items
	report["cart_id"] = _cart.ID
	session.Save()
	c.HTML(http.StatusOK, "cart", report)
}

// 空购物车
func ClearCartHandler(c *gin.Context) {
	session := sessions.Default(c)
	cart := session.Get("cart")
	if cart == nil {
		c.HTML(http.StatusBadRequest, "cart", gin.H{"error": "Cart not found"})
	}
	request := cart.(model.Cart)
	err := service.ClearCart(request.ID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "cart", gin.H{"error": "Failed to clear cart"})
		return
	}

	c.Redirect(http.StatusFound, "/")
}
