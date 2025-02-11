package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProduct .
// @router /product [GET]
func GetProduct(c *gin.Context) {
	c.HTML(http.StatusOK, "product", gin.H{})
}

// SearchProducs .
// @router /search [GET]
func SearchProducs(c *gin.Context) {
	c.HTML(http.StatusOK, "search", gin.H{})
}
