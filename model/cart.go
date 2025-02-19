package model

// CartItem 购物车中的单个商品
type CartItem struct {
	ID       int     `json:"id"`       // 商品ID
	Name     string  `json:"name"`     // 商品名称
	Price    float64 `json:"price"`    // 商品价格
	Quantity int     `json:"quantity"` // 数量
}

// Cart 购物车
type Cart struct {
	UserID int        `json:"user_id"` // 用户ID
	Items  []CartItem `json:"items"`   // 商品列表
}
