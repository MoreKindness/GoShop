package main

import (
	"gomall/dal"
	"gomall/dal/mysql"
	"gomall/router"
	"gomall/service"

	"encoding/gob"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gomall/model"
)

func main() {
	mysql.Init()
	r := gin.Default()
	gob.Register(model.User{})
	gob.Register(model.Cart{})
	gob.Register(model.CartItem{})
	//初始化数据库连接
	mysql.Init()

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
	r.GET("sign-in", func(c *gin.Context) {
		c.HTML(200, "sign-in", gin.H{
			"title": "登录",
			"next":  c.Query("next"),
		})
	})
	r.GET("sign-up", func(c *gin.Context) {
		c.HTML(200, "sign-up", gin.H{
			"title": "注册",
		})
	})
	r.GET("/redirect", func(c *gin.Context) {
		c.HTML(200, "about", gin.H{
			"title": "Error",
		})
	})
	// 调用自动迁移表结构的函数
	dal.MigrateOrderTables()

	router.GeneratedRegister(r)

	// 启动订单取消定时任务并传递数据库连接对象
	go service.CancelExpiredOrders()

	r.Run()
}
