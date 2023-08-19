package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          string     `gorm:"primaryKey;uniqueIndex" json:"productId"`
	Name        string     `json:"name"`
	Description string     `gorm:"size:255" json:"description"`
	Quantity    int        `json:"quantity"`
	Category    []Category `gorm:"foreignKey:ProductID;references:ID" json:"category"`
}
