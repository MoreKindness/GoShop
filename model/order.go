package model

import (
	"time"
)

//type OrderItem struct {
//	ProductId int64 `json:"product_id"`
//	Quantity  int64 `json:"quantity"`
//}
//
//type Order struct {
//	OrderId    int64       `json:"id"`
//	OrderItems []OrderItem `json:"order_items"`
//}

type OrderState string

const (
	OrderStatePlaced   OrderState = "placed"
	OrderStatePaid     OrderState = "paid"
	OrderStateCanceled OrderState = "canceled"
)

type Base struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Consignee 收件人 结构体不单独建表，字段在 Order 表中存储
type Consignee struct {
	Email         string `json:"email"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"status"`
	Country       string `json:"country"`
	ZipCode       int32  `json:"zip_code"` //邮编
}

// OrderItem 一个订单中的单个商品项
type OrderItem struct {
	Base
	ProductId    uint32  `json:"product_id"`
	OrderIdRefer string  `json:"order_id_refer" gorm:"size:256;index"`
	Quantity     int32   `json:"quantity"` //商品数量
	Price        float64 `json:"price"`    //单价
	ProductName  string  `json:"product_name"`
	Picture      string  `json:"picture"`
	Description  string  `json:"description"`
	Cost         float32 `json:"cost"` //成本
}

// Order 一个订单，包含订单的基本信息
type Order struct {
	Base
	OrderId      string      `json:"order_id" gorm:"uniqueIndex;size:256"`
	UserId       uint32      `json:"user_id"`
	UserCurrency string      `json:"user_currency"`             // 用户使用的货币
	Consignee    Consignee   `json:"consignee" gorm:"embedded"` //gorm:"embedded" :将一个结构体嵌入到另一个结构体中,Order.Consignee访问
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
	OrderState   OrderState  `json:"order_state"`
}
