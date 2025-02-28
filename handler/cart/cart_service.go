package cart

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	OrderId     string
	CreatedDate string
	Items       []Item
}

type Item struct {
	Picture     string
	ProductName string
	Qty         int
	Cost        float64
}

// AddCartItem .
// @router /cart [POST]
func AddCartItem(c *gin.Context) {
	c.HTML(http.StatusOK, "cart", gin.H{})
}

// GetCart .
// @router /cart [GET]
func GetCart(c *gin.Context) {
	fmt.Println("GetCart")
	orders := []Order{
		{
			OrderId:     "123",
			CreatedDate: "2022-01-01",
			Items: []Item{
				{
					Picture:     "image1.jpg",
					ProductName: "Product 1",
					Qty:         2,
					Cost:        10.0,
				},
				{
					Picture:     "image2.jpg",
					ProductName: "Product 2",
					Qty:         1,
					Cost:        20.0,
				},
			},
		},
	}
	c.HTML(http.StatusOK, "cart", gin.H{"orders": orders})
}
