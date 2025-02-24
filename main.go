package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gomall/dal"
	"gomall/dal/dao"
	"gomall/dal/sql"
	"gomall/handler/auth"
	"gomall/handler/auth/middleware"
	"gomall/service"
	"gorm.io/gorm"
)

func main() {

	db := sql.Init()

	server := initWebServer()
	initUserHdl(db, server)

	//
	//r.GET("/login", func(ctx *gin.Context) {
	//	ctx.HTML(http.StatusOK, "sign-in", gin.H{"hello": "ggg"})
	//})

	//router.GeneratedRegister(r)
	server.Run(":8012")
}

func initUserHdl(db *gorm.DB, server *gin.Engine) {

	ud := dao.NewUserDAO(db)
	ur := dal.NewUserRepository(ud)
	us := service.NewUserService(ur)
	hdl := auth.NewUserHandler(us)
	hdl.RegisterRoutes(server)
}

func initWebServer() *gin.Engine {
	server := gin.Default()

	server.LoadHTMLGlob("template/*")
	server.Delims("{{", "}}")
	server.Static("/static", "./static")

	//server.Use(cors.New(cors.Config{}))
	login := &middleware.LoginMiddlewareBuilder{}

	// 直接存 cookie
	store := cookie.NewStore([]byte("goshop"))
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())
	return server
}
