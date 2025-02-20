package service

import (
	"errors"
	"gomall/model"
	"sync"
)

// 模拟数据库，使用内存存储购物车数据
var cartStorage = make(map[int]*model.Cart)
var mu sync.Mutex

// AddToCart 添加商品到购物车
func AddToCart(userID int, item model.CartItem) error {
	mu.Lock()
	defer mu.Unlock()

	// 如果用户购物车不存在，初始化一个
	if _, exists := cartStorage[userID]; !exists {
		cartStorage[userID] = &model.Cart{
			UserID: userID,
			Items:  []model.CartItem{},
		}
	}

	// 检查商品是否已存在，如果存在则更新数量
	for i, v := range cartStorage[userID].Items {
		if v.ProductID == item.ProductID {
			cartStorage[userID].Items[i].Quantity += item.Quantity
			return nil
		}
	}

	// 添加新商品
	cartStorage[userID].Items = append(cartStorage[userID].Items, item)
	return nil
}

// GetCart 获取用户购物车
func GetCart(userID int) (*model.Cart, error) {
	mu.Lock()
	defer mu.Unlock()

	cart, exists := cartStorage[userID]
	if !exists {
		return nil, errors.New("购物车为空")
	}
	return cart, nil
}

// ClearCart 清空购物车
func ClearCart(userID int) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := cartStorage[userID]; !exists {
		return errors.New("购物车已为空")
	}

	delete(cartStorage, userID)
	return nil
}
