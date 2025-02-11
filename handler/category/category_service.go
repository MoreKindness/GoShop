package category

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Category .
// @router /category/:category [GET]
func Category(c *gin.Context) {
	c.HTML(http.StatusOK, "category", gin.H{})
}
