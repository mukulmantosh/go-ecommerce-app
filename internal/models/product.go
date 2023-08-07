package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductID        string           `gorm:"primaryKey;uniqueIndex" json:"productId"`
	Name             string           `json:"name"`
	Description      string           `gorm:"size:255" json:"description"`
	Category         Category         `gorm:"foreignKey:CategoryID;references:ProductID" json:"category,omitempty"`
	ProductInventory ProductInventory `gorm:"foreignKey:InventoryID;references:ProductID" json:"product_inventory,omitempty"`
}

type ProductInventory struct {
	InventoryID string `gorm:"primaryKey;uniqueIndex" json:"inventoryId"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
}
