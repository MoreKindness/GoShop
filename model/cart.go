package model

// CartItem 购物车中的单个商品
//type CartItem struct {
//	ID        uint    `json:"id" gorm:"primaryKey"`
//	ProductID string  `json:"product_id"` // 商品ID
//	Name      string  `json:"name"`       // 商品名称
//	Price     float64 `json:"price"`      // 商品价格
//	Quantity  int     `json:"quantity"`   // 数量
//}

// Cart 购物车
type Cart struct {
	ID    uint       `json:"id" gorm:"primaryKey"`
	Items []CartItem `json:"items" gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	CartID      uint    `json:"cart_id"`
	ProductID   uint    `json:"product_id"`
	Name        string  `json:"name"`
	Picture     string  `json:"picture"`
	Description string  `json:"description"`
	Quantity    uint    `json:"quantity"`
	Price       float64 `json:"price"`
}

// // 购物车项模型
// type CartItem struct {
// 	ID        uint    `gorm:"primaryKey"`
// 	UserID    int     `gorm:"not null"`  // 用户ID
// 	ProductID int     `gorm:"not null"`  // 商品ID
// 	Quantity  int     `gorm:"default:1"` // 商品数量
// 	Price     float64 `gorm:"not null"`  // 商品价格
// 	CreatedAt string  `gorm:"not null"`
// 	UpdatedAt string  `gorm:"not null"`
// }
