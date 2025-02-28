package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Id      int
	Name    string
	Price   float64
	Picture string
}

// Home .
// @router / [GET]
func Home(c *gin.Context) {
	items := []Item{
		{Id: 1, Name: "Product 1", Price: 100.0, Picture: "static/t-shit-1.jpg"},
		{Id: 2, Name: "Product 2", Price: 200.0, Picture: "t-shit-1.jpg"},
		{Id: 3, Name: "Product 3", Price: 300.0, Picture: "t-shit-1.jpg"},
	}
	c.HTML(http.StatusOK, "home", gin.H{
		"user_id":  "1234567890",
		"cart_num": "10",
		"items":    items,
	})
}
