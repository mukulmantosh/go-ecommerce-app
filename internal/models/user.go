package models

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mukulmantosh/go-ecommerce-app/internal/utils"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	ID        string        `gorm:"primaryKey" json:"ID"`
	Username  string        `gorm:"uniqueIndex;size:255" json:"username"`
	Password  string        `json:"password"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Address   []UserAddress `gorm:"foreignKey:UserID;references:ID" json:"address"`
	Cart      Cart          `gorm:"foreignKey:UserID;references:ID" json:"cart"`
}

type UserAddress struct {
	ID         string `gorm:"primaryKey" json:"userAddressId"`
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	Mobile     string `json:"mobile"`
	UserID     string `json:"user_id"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomJWTClaims struct {
	Name     string `json:"name"`
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
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
	fmt.Println("password ", u.Password)
	hash, err := utils.GenerateFromPassword(u.Password, params)
	if err != nil {
		log.Fatal(err)
	}
	tx.Model(u).Update("Password", hash)
	return
}

func (l Login) VerifyPassword(hashPassword string) (match bool, err error) {
	fmt.Println(l.Password, hashPassword)
	match, err = utils.ComparePasswordAndHash(l.Password, hashPassword)
	if err != nil {
		return false, err
	}
	return match, err
}
