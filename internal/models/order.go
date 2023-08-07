package models

import "gorm.io/gorm"

type OrderSession struct {
	gorm.Model
	OrderSessionID string `gorm:"primaryKey;uniqueIndex" json:"orderSessionId"`
	total          float32
}
