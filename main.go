package main

import (
	"gomall/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//goshop 是用于加密的参数，可以随意设置
	store := cookie.NewStore([]byte("goshop"))
	r.Use(sessions.Sessions("goshop", store))
	r.LoadHTMLGlob("template/*")
	r.Delims("{{", "}}")
	r.Static("/static", "./static")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})
	router.GeneratedRegister(r)
	r.Run()
}
