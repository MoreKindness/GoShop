package model

import (
	"time"
)

type Cart struct {
	ID        uint      `gorm:"primary_key"`
	UserID    uint      `gorm:"index"`
	Products  []Product `gorm:"many2many:cart_products;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	ID       uint   `gorm:"primary_key"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
