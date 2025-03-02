package service

import (
	"gomall/dal"
	"gomall/dal/mysql"
	"gomall/model"
)

// 获取购物车信息
func GetCart(cartID uint) (*model.Cart, error) {
	return dal.GetCart(mysql.DB, cartID)
}

// 添加商品到购物车
func AddToCart(cartID uint, item model.CartItem) error {
	return dal.AddToCart(mysql.DB, cartID, item)
}

// 清空购物车
func ClearCart(cartID uint) error {
	return dal.ClearCart(mysql.DB, cartID) // 直接调用 DAL 层方法
}

// func ClearCart(cartID uint) error {
// 	// 删除 cart_items 表中所有 cart_id = cartID 的商品
// 	if err := mysql.DB.Where("cart_id = ?", cartID).Delete(&model.CartItem{}).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// 创建购物车
func CreateCart() (*model.Cart, error) {
	cart := &model.Cart{}
	if err := mysql.DB.Create(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

// func CheckDBConnection() error {
// 	db, err := mysql.DB.DB() // 获取底层 SQL 连接
// 	if err != nil {
// 		return err
// 	}
// 	return db.Ping() // 发送 Ping 以检查数据库是否正常连接
// }
