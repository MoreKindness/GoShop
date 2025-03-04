package service

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"gomall/dal"
	"gomall/dal/mysql"
	"gomall/model"
	"log"
	"strconv"
	"time"
)

// PlaceOrder 处理创建订单的请求
func PlaceOrder(order *model.Order) error {
	// 在创建订单对象之前生成一个唯一的 Id
	//orderId := uuid.New().String()
	node, _ := snowflake.NewNode(1)
	Id := node.Generate()
	orderId := strconv.FormatUint(uint64(Id), 10)

	base := model.Base{
		ID:        uint(Id),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	o := &model.Order{
		Base:         base,
		UserId:       order.UserId,
		OrderId:      orderId, // 使用 UUID 作为 OrderId
		UserCurrency: order.UserCurrency,
		Consignee:    order.Consignee,
		OrderItems:   make([]model.OrderItem, len(order.OrderItems)),
	}

	for i, item := range order.OrderItems {
		o.OrderItems[i] = model.OrderItem{
			ProductId:    item.ProductId,
			OrderIdRefer: orderId, // 使用相同的 OrderId
			Quantity:     item.Quantity,
			Price:        item.Price,
			ProductName:  item.ProductName,
			Picture:      item.Picture,
		}
	}

	if err := dal.CreateOrder(mysql.DB, o); err != nil {
		return err
	}

	return nil
}

// UpdateOrder 更新订单信息
func UpdateOrder(ID uint, userID uint32, consignee *model.Consignee, updatedAt time.Time) error {
	order, err := dal.GetOrder(mysql.DB, ID)
	if err != nil {
		return fmt.Errorf("获取订单失败: %w", err)
	}

	if order.UserId != userID {
		return fmt.Errorf("用户无权修改该订单")
	}

	order.Consignee = *consignee
	order.UpdatedAt = updatedAt

	if err := dal.UpdateOrder(mysql.DB, order); err != nil {
		return fmt.Errorf("更新订单失败: %w", err)
	}

	return nil
}

// ShowOrder 显示某个订单
func ShowOrder(ID uint) (*model.Order, error) {
	order, err := dal.GetOrder(mysql.DB, ID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// GetAllOrders 获取用户的所有订单
func GetAllOrders(userID uint32) ([]model.Order, error) {
	orders, err := dal.ListOrders(mysql.DB, userID)
	if err != nil {
		return nil, fmt.Errorf("获取用户所有订单失败: %w", err)
	}
	return orders, nil

}

// CancelOrder 取消订单
func CancelOrder(ID uint) error {
	if err := dal.CancelOrder(mysql.DB, ID); err != nil {
		return fmt.Errorf("取消订单失败: %w", err)
	}
	return nil
}

// // CancelExpiredOrders 定时任务：取消过期未支付的订单
func CancelExpiredOrders() {
	// 每5分钟检查一次未支付的订单
	var minutes int = 5
	ticker := time.NewTicker(time.Duration(minutes) * time.Minute)
	for {
		select {
		case <-ticker.C:
			orders, err := dal.GetUnpaidOrders(mysql.DB, minutes)
			if err != nil {
				log.Printf("获取相应订单失败: %v", err)
				continue
			}

			// 遍历未支付的订单，并进行取消操作
			for _, order := range orders {
				err := dal.UpdateOrderStatus(mysql.DB, order.ID)
				if err != nil {
					log.Printf("取消订单 %s 失败: %v", order.ID, err)
					continue
				}
				log.Printf("订单 %s 已成功取消", order.ID)
			}
		}
	}
}
