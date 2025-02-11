package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddCartItem .
// @router /cart [POST]
func AddCartItem(c *gin.Context) {
	c.HTML(http.StatusOK, "cart", gin.H{})
}

// GetCart .
// @router /cart [GET]
func GetCart(c *gin.Context) {
	c.HTML(http.StatusOK, "cart", gin.H{})
}
