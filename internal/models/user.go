package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string        `gorm:"primaryKey;uniqueIndex" json:"ID"`
	Username  string        `gorm:"uniqueIndex;size:255" json:"username"`
	Password  string        `json:"password"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Address   []UserAddress `gorm:"foreignKey:UserID;references:ID" json:"address"`
}

type UserAddress struct {
	ID         string `gorm:"primaryKey;uniqueIndex" json:"userAddressId"`
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	Mobile     string `json:"mobile"`
	UserID     string `json:"user_id"`
}
