package about

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func About(c *gin.Context) {
	user_id, err := c.Cookie("user_id")
	if err != nil {
		c.String(http.StatusForbidden, err.Error())
	}
	cart_data, err := c.Cookie("cart_data")
	if err != nil {
		c.String(http.StatusForbidden, err.Error())
	}
	c.HTML(http.StatusOK, "about", gin.H{"user_id": user_id, "cart_data": cart_data})
}
