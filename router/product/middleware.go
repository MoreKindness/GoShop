package product

import (
	"github.com/gin-gonic/gin"
)

func RootMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(c *gin.Context) { // your code...
		},
	}
}

func GetproductMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(c *gin.Context) { // your code...
		},
	}
}

func SearchProductsMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		// 参数校验中间件
		func(c *gin.Context) {
		},
	}
}
