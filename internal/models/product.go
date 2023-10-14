package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          string  `gorm:"primaryKey" json:"ID"`
	Name        string  `json:"name"`
	Description string  `gorm:"size:255" json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CategoryID  string  `json:"category_id"`
}

type ProductParams struct {
	ProductID string `json:"productID"`
}
