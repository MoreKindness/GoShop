package category

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gomall/model"
	"gomall/service"
	"net/http"
	"reflect"
)

// Category .
// @router /category/:category [GET]
func Category(c *gin.Context) {
	category := c.Param("category")
	session := sessions.Default(c)
	report := make(map[string]interface{})
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
	product_service := service.NewProductService()
	products, err := product_service.ListProductsByCategory(category, 1, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		report["items"] = products
	}
	fmt.Println(category, reflect.TypeOf(category))
	c.HTML(http.StatusOK, "category", report)
}
