package main

import (
	"gomall/dal"
	"gomall/dal/mysql"
	"gomall/router"
	"gomall/service"

	"gomall/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//初始化数据库连接
	mysql.Init()

	store := cookie.NewStore([]byte("goshop"))
	r.Use(sessions.Sessions("goshop", store))
	r.LoadHTMLGlob("template/*")
	r.Delims("{{", "}}")
	r.Static("/static", "./static")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})

	// 调用自动迁移表结构的函数
	dal.MigrateOrderTables()

	router.GeneratedRegister(r)

	// 启动订单取消定时任务并传递数据库连接对象
	go service.CancelExpiredOrders()

	r.Run()
}
