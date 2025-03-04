package dal

import (
	"gomall/dal/mysql"
	"gomall/model"

	"gorm.io/gorm"
)

// 自动迁移表结构
func MigrateCartTables() {
	mysql.DB.AutoMigrate(&model.Cart{}, &model.CartItem{})
}

// 获取购物车信息
func GetCart(db *gorm.DB, cartID uint) (*model.Cart, error) {
	var cart model.Cart
	if err := db.Preload("Items").First(&cart, cartID).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

// 添加商品到购物车
func AddToCart(db *gorm.DB, cartID uint, item model.CartItem) (model.CartItem, error) {
	var existingItem model.CartItem

	if err := db.Where("cart_id = ? AND product_id = ?", cartID, item.ProductID).
		First(&existingItem).Error; err == nil {
		existingItem.Quantity += item.Quantity
		return existingItem, db.Save(&existingItem).Error
	}

	item.CartID = cartID
	return item, db.Create(&item).Error
}

// 清空购物车
func ClearCart(db *gorm.DB, cartID uint) error {
	// 删除 cart_items 表中所有 cart_id = cartID 的商品
	if err := db.Where("cart_id = ?", cartID).Delete(&model.CartItem{}).Error; err != nil {
		return err
	}
	return nil
}

// func ClearCart(db *gorm.DB, cartID uint) error {
// 	return db.Where("cart_id = ?", cartID).Delete(&model.CartItem{}).Error
// }
