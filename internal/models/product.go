package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          string `gorm:"primaryKey;uniqueIndex" json:"ID"`
	Name        string `json:"name"`
	Description string `gorm:"size:255" json:"description"`
	Quantity    int    `json:"quantity"`
	CategoryID  string `json:"category_id"`
}

//type User struct {
//	gorm.Model
//	MemberNumber string
//	CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number     string
//	UserNumber string
//}
