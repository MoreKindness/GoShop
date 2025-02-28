package model

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	Consignee   Consignee
	OrderId     string
	CreatedDate string
	OrderState  string
	Cost        float32
	Items       []Product
}
