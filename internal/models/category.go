package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          string    `gorm:"primaryKey;uniqueIndex" json:"ID"`
	Name        string    `gorm:"uniqueIndex;size:255" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	Products    []Product `gorm:"foreignKey:CategoryID;references:ID" json:"product"`
}
