package order

import (
	"gorm.io/gorm"
)

type Product struct {
	ID          int
	Name        string
	Picture     string
	Description string
	Price       float64
}

type Order struct {
	gorm.Model
	UserID      uint
	CheckoutID  uint
	OrderId     string
	CreatedDate string
	OrderState  string
	Cost        float32
	Items       []Product
}

func NewOrder() *Order {
	return &Order{}
}

func (Order) TableName() string {
	return "orders"
}

func (o *Order) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

func (o *Order) Update(db *gorm.DB) error {
	return db.Save(o).Error
}

func (o *Order) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

func (o *Order) Get(db *gorm.DB, id uint) error {
	return db.First(o, id).Error
}

func (o *Order) GetByOrderId(db *gorm.DB, orderId string) error {
	return db.Where("order_id = ?", orderId).First(o).Error
}

func (o *Order) GetByCheckoutId(db *gorm.DB, checkoutId uint) error {
	return db.Where("checkout_id = ?", checkoutId).First(o).Error
}

func (o *Order) GetByUserId(db *gorm.DB, userId uint) ([]Order, error) {
	//返回根据用户id查到的所有订单
	var orders []Order
	error := db.Where("user_id = ?", userId).Find(&orders).Error
	return orders, error
}

func (o *Order) GetAll(db *gorm.DB) ([]Order, error) {
	var orders []Order
	return orders, db.Find(&orders).Error
}
