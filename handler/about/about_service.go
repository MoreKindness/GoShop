package about

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about", gin.H{"user_id": c.Get("user_id")})
}
