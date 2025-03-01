package main

import (
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
	//goshop 是用于加密的参数，可以随意设置
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
		c.HTML(200, "about", gin.H{})
	})
	router.GeneratedRegister(r)
	r.Run()
}
