package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Home .
// @router / [GET]
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{})
}
