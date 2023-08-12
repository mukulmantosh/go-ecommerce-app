package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          string `gorm:"primaryKey;uniqueIndex" json:"categoryId"`
	Name        string `gorm:"uniqueIndex;size:255" json:"name"`
	Description string `gorm:"size:255" json:"description"`
	ProductID   string `json:"product_id"`
}
