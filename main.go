package main

import (
	"gomall/dal"
	"gomall/dal/mysql"
	"gomall/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gomall/router"
)

func main() {
	r := gin.Default()
	//goshop 是用于加密的参数，可以随意设置
	store := cookie.NewStore([]byte("goshop"))
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())
	return server
}
