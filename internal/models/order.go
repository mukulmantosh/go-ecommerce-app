package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID       string    `gorm:"primaryKey;uniqueIndex" json:"cartId"`
	Total    float32   `json:"total"`
	UserID   string    `json:"user_id"`
	Products []Product `gorm:"many2many:cart_items;foreignKey:ID;joinForeignKey:ProductID;References:ID;joinReferences:CartID"`
}
