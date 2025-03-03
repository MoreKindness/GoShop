package dal

import (
	"fmt"
	"gomall/dal/mysql"
	"gomall/model"
	"gorm.io/gorm"
	"log"
	"time"
)

// MigrateOrderTables 自动迁移订单表结构
func MigrateOrderTables() {
	err := mysql.DB.AutoMigrate(&model.Order{}, &model.OrderItem{})
	if err != nil {
		log.Println("failed to migrate database: ", err)
	}
}

// CreateOrder 创建新的订单
func CreateOrder(db *gorm.DB, order *model.Order) error {
	return db.Create(order).Error
}

// UpdateOrder 更新现有的订单信息
func UpdateOrder(db *gorm.DB, order *model.Order) error {
	return db.Save(order).Error
}

// GetOrder 根据订单ID获取订单信息
func GetOrder(db *gorm.DB, ID uint) (*model.Order, error) {
	var order model.Order
	if err := db.Preload("OrderItems").First(&order, "id = ?", ID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// ListOrders 获取用户的所有订单
func ListOrders(db *gorm.DB, userID uint32) ([]model.Order, error) {
	var orders []model.Order
	if err := db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// CancelOrder 取消未支付未过期订单
func CancelOrder(db *gorm.DB, ID uint) error {
	var order model.Order
	if err := db.Where("id = ?", ID).First(&order).Error; err != nil {
		return fmt.Errorf("订单未找到: %w", err)
	}

	// 如果已支付，则无法取消
	if order.OrderState == model.OrderStatePaid {
		return fmt.Errorf("订单已支付，无法取消")
	}

	// 如果订单已经取消，无法再次取消
	if order.OrderState == model.OrderStateCanceled {
		return fmt.Errorf("订单已经取消，无需再次取消")
	}

	if err := db.Where("order_id = ?", ID).First(&order).Error; err != nil {
		return err
	}
	order.OrderState = model.OrderStateCanceled
	return UpdateOrderStatus(db, order.ID)
}

// GetUnpaidOrders 获取未支付的订单，且创建时间超过指定的分钟数
// minutes 自定义的时间间隔（单位：分钟）
// createdAt 查询基准时间
func GetUnpaidOrders(db *gorm.DB, minutes int) ([]model.Order, error) {
	var orders []model.Order
	var validOrders []model.Order // 用于存储符合条件的订单

	// 查询未支付的订单
	err := db.Where("order_state = ?", model.OrderStatePlaced).Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("查询未支付订单失败: %v", err)
	}
	//查询未支付的订单，且创建时间超过传入的时间
	for _, o := range orders {
		// 计算时间差，使用传入的 createdAt 和时间间隔来查询
		thresholdTime := o.CreatedAt.Add(time.Duration(-minutes) * time.Minute)
		if db.Where("created_at <= ?", thresholdTime, model.OrderStatePlaced).Find(&o) != nil {
			validOrders = append(validOrders, o) // 将符合条件的订单添加到新的切片中
		}
	}
	if validOrders == nil {
		return nil, fmt.Errorf("查询未支付且创建时间超过指定分钟数的订单失败: %v", err)
	}
	return validOrders, nil
}

// UpdateOrderStatus 更新未支付过期订单状态为已取消
func UpdateOrderStatus(db *gorm.DB, ID uint) error {

	// 更新订单状态为已取消
	if err := db.Model(&model.Order{}).Where("id = ?", ID).Update("status", model.OrderStateCanceled).Error; err != nil {
		return err
	}
	return nil
}

//// ScheduleOrderCancellation 定时取消订单
//func ScheduleOrderCancellation(db *gorm.DB, minutes int, createdAt time.Time,ID uint) error {
//	// 检查订单是否创建超过 minutes 分钟
//	if time.Since(createdAt) > time.Duration(minutes)*time.Minute {
//		return UpdateOrderStatus(db, ID)
//	}
//	return nil
//}
