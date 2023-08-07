package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID       string       `gorm:"primaryKey;uniqueIndex" json:"userId"`
	Username     string       `gorm:"uniqueIndex;size:255" json:"username"`
	Password     string       `json:"password"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	UserAddress  UserAddress  `gorm:"foreignKey:UserAddressID;references:UserID" json:"address,omitempty"`
	OrderSession OrderSession `gorm:"foreignKey:OrderSessionID;references:UserID" json:"order_session,omitempty"`
}

type UserAddress struct {
	UserAddressID string `gorm:"primaryKey;uniqueIndex" json:"UserAddressId"`
	Address       string `json:"address"`
	City          string `json:"city"`
	PostalCode    string `json:"postal_code"`
	Country       string `json:"country"`
	Mobile        string `json:"mobile"`
}
