package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID       string    `gorm:"primaryKey" json:"cartId"`
	UserID   string    `json:"user_id"`
	Products []Product `gorm:"many2many:cart_items;"`
}

type Order struct {
	gorm.Model
	ID          string    `gorm:"primaryKey" json:"orderId"`
	Total       float32   `gorm:"default:0" json:"total"`
	UserID      string    `json:"user_id"`
	OrderStatus string    `json:"order_status"`
	Products    []Product `gorm:"many2many:order_details;"`
}
