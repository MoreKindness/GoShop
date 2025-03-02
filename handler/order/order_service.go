package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderList .
// @router /order [GET]
func OrderList(c *gin.Context) {
	// user_id, err := c.Cookie("user_id")
	// if err != nil {
	// 	c.String(http.StatusForbidden, err.Error())
	// }
	// cart_data, err := c.Cookie("cart_data")
	// if err != nil {
	// 	c.String(http.StatusForbidden, err.Error())
	// }
	c.HTML(http.StatusOK, "order", gin.H{})
}
