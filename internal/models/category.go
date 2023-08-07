package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryID  string `gorm:"primaryKey;uniqueIndex" json:"categoryId"`
	Name        string `json:"name"`
	Description string `gorm:"size:255" json:"description"`
}
