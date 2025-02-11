package checkout

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Checkout .
// @router /checkout [GET]
func Checkout(c *gin.Context) {
	c.HTML(http.StatusOK, "checkout", gin.H{})
}

// CheckoutWaiting .
// @router /checkout/waiting [POST]
func CheckoutWaiting(c *gin.Context) {
	c.HTML(http.StatusOK, "waiting", gin.H{})
}

// CheckoutResult .
// @router /checkout/result [GET]
func CheckoutResult(c *gin.Context) {
	c.HTML(http.StatusOK, "result", gin.H{})
}
