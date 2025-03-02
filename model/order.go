package model

type OrderItem struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

type Order struct {
	OrderId    int64       `json:"id"`
	OrderItems []OrderItem `json:"order_items"`
}
