package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register .
// @router /auth/register [POST]
func Register(c *gin.Context) {
	//session := sessions.Default(c)
	//if session.Get("user_id") == nil {
	//	c.HTML(http.StatusOK, "sign_up", gin.H{})
	//}
	c.Redirect(http.StatusFound, "/")
}

// Login .
// @router /auth/login [POST]
func Login(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}

// Logout .
// @router /auth/logout [POST]
func Logout(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
