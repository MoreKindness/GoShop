package service

import (
	"gomall/dal/mysql"
	"gomall/dal/order"
)

type ServiceOrder struct {
	OrderID    string          `json:"order_id"`
	CreateDate string          `json:"create_data"`
	OrderState string          `json:"order_state"`
	Items      []order.Product `json:"items"`
}

func GetOrderListByUserID(userID uint) ([]ServiceOrder, error) {
	var orderDB = order.NewOrder()

	// TODO: 获取订单列表
	orderList, err := orderDB.GetByUserId(mysql.DB, userID)
	if err != nil {
		return nil, err
	}

	var serviceOrderList []ServiceOrder
	for _, order := range orderList {
		serviceOrderList = append(serviceOrderList, ServiceOrder{
			OrderID:    order.OrderId,
			CreateDate: order.CreatedDate,
			OrderState: order.OrderState,
			Items:      order.Items,
		})
	}
	return serviceOrderList, nil
}
