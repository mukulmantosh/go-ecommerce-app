package models

import (
	"github.com/mukulmantosh/go-ecommerce-app/internal/utils"
	"gorm.io/gorm"
	"log"
)

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

// AfterCreate Password Hashing
func (u User) AfterCreate(tx *gorm.DB) (err error) {

	params := &utils.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := utils.GenerateFromPassword(u.Password, params)
	if err != nil {
		log.Fatal(err)
	}
	tx.Model(u).Update("Password", hash)
	return
}
