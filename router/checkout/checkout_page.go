package checkout

import (
	"github.com/gin-gonic/gin"
	"gomall/handler/checkout"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *gin.Engine) {
	root := r.Group("/", RootMw()...)
	root.GET("/checkout", append(Checkout0Mw(), checkout.Checkout)...)
	_checkout := root.Group("/checkout", CheckoutMw()...)
	_checkout.GET("/result", append(CheckoutresultMw(), checkout.CheckoutResult)...)
	_checkout.POST("/waiting", append(CheckoutwaitingMw(), checkout.CheckoutWaiting)...)
}
