package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// OrderList .
// @router /order [GET]
func OrderList(c *gin.Context) {
	c.HTML(http.StatusOK, "order", gin.H{})
}
