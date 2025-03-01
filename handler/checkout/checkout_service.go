package checkout

import (
	"fmt"
	"gomall/dal/mysql"
	"gomall/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CheckoutForm struct {
	Email           string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" form:"email"`
	Firstname       string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty" form:"firstname"`
	Lastname        string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty" form:"lastname"`
	Street          string `protobuf:"bytes,4,opt,name=street,proto3" json:"street,omitempty" form:"street"`
	Zipcode         string `protobuf:"bytes,5,opt,name=zipcode,proto3" json:"zipcode,omitempty" form:"zipcode"`
	Province        string `protobuf:"bytes,6,opt,name=province,proto3" json:"province,omitempty" form:"province"`
	Country         string `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty" form:"country"`
	City            string `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty" form:"city"`
	CardNum         string `protobuf:"bytes,9,opt,name=card_num,json=cardNum,proto3" json:"card_num,omitempty" form:"cardNum"`
	ExpirationMonth int32  `protobuf:"varint,10,opt,name=expiration_month,json=expirationMonth,proto3" json:"expiration_month,omitempty" form:"expirationMonth"`
	ExpirationYear  int32  `protobuf:"varint,11,opt,name=expiration_year,json=expirationYear,proto3" json:"expiration_year,omitempty" form:"expirationYear"`
	Cvv             int32  `protobuf:"varint,12,opt,name=cvv,proto3" json:"cvv,omitempty" form:"cvv"`
	Payment         string `protobuf:"bytes,13,opt,name=payment,proto3" json:"payment,omitempty" form:"payment"`
}

// Checkout .
// @router /checkout [GET]
func Checkout(c *gin.Context) {
	fmt.Println(mysql.DB)
	session := sessions.Default(c)
	var mycart model.Cart
	mycart.UserID = 123
	session.Set("cart", mycart)
	session.Save()
	var cart = session.Get("cart")
	fmt.Println("cart: ", cart)
	c.HTML(http.StatusOK, "checkout", gin.H{})
}

// CheckoutWaiting .
// @router /checkout/waiting [POST]
func CheckoutWaiting(c *gin.Context) {
	session := sessions.Default(c)
	cart := session.Get("cart")
	user_id := session.Get("user_id")
	if cart == nil {
	}
	fmt.Println("cart: ", cart)
	fmt.Println("user_id: ", user_id)
	var form CheckoutForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println("err: ", err)
		fmt.Println("错了")
	}
	fmt.Println("Email:", form.Email)
	fmt.Println("Firstname:", form.Firstname)
	fmt.Println("Lastname:", form.Lastname)
	c.HTML(http.StatusOK, "waiting", gin.H{})
}

// CheckoutResult .
// @router /checkout/result [GET]
func CheckoutResult(c *gin.Context) {
	c.HTML(http.StatusOK, "result", gin.H{})
}
