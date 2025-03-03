package main

import (
	"gomall/dal"
	"gomall/dal/mysql"
	"gomall/router"
	"gomall/service"

	"encoding/gob"
	"fmt"
	"gomall/dal/mysql"
	"gomall/model"
	"gomall/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	mysql.Init()
	fmt.Println(mysql.DB)
	r := gin.Default()
	gob.Register(model.User{})
	gob.Register(model.Cart{})
	gob.Register(model.CartItem{})

	//初始化数据库连接
	mysql.Init()

	// 调用自动迁移表结构的函数
	dal.MigrateOrderTables()

	store := cookie.NewStore([]byte("goshop"))
	r.Use(sessions.Sessions("goshop", store))

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(cors.Default())

	r.LoadHTMLGlob("template/*")
	r.Delims("{{", "}}")
	r.Static("/static", "./static")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})

	router.GeneratedRegister(r)

	// 启动订单取消定时任务并传递数据库连接对象
	go service.CancelExpiredOrders()

	r.Run(":8080")
}
