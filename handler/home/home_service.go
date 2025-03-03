package home

import (
	"gomall/model"
	"gomall/service"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Item struct {
	Id      int
	Name    string
	Price   float64
	Picture string
}

// Home .
// @router / [GET]
func Home(c *gin.Context) {
	var report = make(map[string]interface{})
	session := sessions.Default(c)
	user_id := session.Get("user_id")
	if user_id != nil {
		report["user_id"] = user_id
	}
	cart := session.Get("cart")
	var _cart model.Cart
	if cart != nil {
		_cart = cart.(model.Cart)
		report["cart_num"] = len(_cart.Items)
	}
	session.Save()
	product_service := service.NewProductService()
	products, err := product_service.ListProducts(1, 4)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	report["items"] = products
	c.HTML(http.StatusOK, "home", report)
}
