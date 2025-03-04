package checkout

import (
	"gomall/model"
	"gomall/service"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CheckoutForm struct {
	Email           string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" form:"email"`
	Firstname       string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty" form:"firstname"`
	Lastname        string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty" form:"lastname"`
	Street          string `protobuf:"bytes,4,opt,name=street,proto3" json:"street,omitempty" form:"street"`
	Zipcode         int32  `protobuf:"bytes,5,opt,name=zipcode,proto3" json:"zipcode,omitempty" form:"zipcode"`
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
	report := make(map[string]interface{})
	session := sessions.Default(c)
	var _cart model.Cart
	var cart = session.Get("cart")
	if cart == nil {
		c.HTML(http.StatusBadRequest, "home", gin.H{"error": "用户登录已失效"})
	} else {
		//转换cart格式
		_cart = cart.(model.Cart)
		report["cart_num"] = len(_cart.Items)
	}
	user := session.Get("user_id")
	if user != nil {
		report["user_id"] = user
	}
	var total float64 = 0
	for _, v := range _cart.Items {
		total += v.Price * float64(v.Quantity)
	}
	report["total"] = total
	c.HTML(http.StatusOK, "checkout", report)
}

// CheckoutWaiting .
// @router /checkout/waiting [POST]
func CheckoutWaiting(c *gin.Context) {
	session := sessions.Default(c)
	cart := session.Get("cart")
	user_id := session.Get("user_id")
	if cart == nil {
		c.HTML(http.StatusOK, "waiting", gin.H{
			"error": "未登录用户",
		})
		return
	}
	_cart := cart.(model.Cart)
	var form CheckoutForm
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusOK, "waiting", gin.H{
			"cart":    cart,
			"user_id": user_id,
			"error":   err.Error(),
		})
		return
	}
	order := model.Order{
		UserId:       uint32(_cart.ID),
		UserCurrency: "$",
		OrderState:   model.OrderStatePaid,
		Consignee: model.Consignee{
			Email:         form.Email,
			StreetAddress: form.Street,
			City:          form.City,
			State:         form.Province,
			Country:       form.Country,
			ZipCode:       form.Zipcode,
		},
	}
	for _, v := range _cart.Items {
		order.OrderItems = append(order.OrderItems, model.OrderItem{
			ProductId:   uint32(v.ProductID),
			Quantity:    int32(v.Quantity),
			Cost:        float32(v.Price),
			Price:       v.Price,
			ProductName: v.Name,
		})
	}
	service.PlaceOrder(&order)
	c.Redirect(http.StatusMovedPermanently, "/checkout/result")
}

// CheckoutResult .
// @router /checkout/result [GET]
func CheckoutResult(c *gin.Context) {
	session := sessions.Default(c)
	user_id := session.Get("user_id")
	cart := session.Get("cart")
	if cart == nil {
		c.HTML(http.StatusOK, "result", gin.H{"error": "用户登录已失效"})
		return
	}
	_cart := cart.(model.Cart)

	c.HTML(http.StatusOK, "result", gin.H{
		"user_id":  user_id,
		"cart_num": len(_cart.Items),
	})
}
