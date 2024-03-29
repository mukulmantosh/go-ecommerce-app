/*
	Copyright 2022 Google LLC

#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
*/
package models

import (
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

type DisplayUser struct {
	ID        string `gorm:"primaryKey" json:"ID"`
	Username  string `gorm:"uniqueIndex;size:255" json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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
	hash, err := utils.GenerateFromPassword(u.Password, params)
	if err != nil {
		log.Fatal(err)
	}
	tx.Model(u).Update("Password", hash)
	return
}

func (l Login) VerifyPassword(hashPassword string) (match bool, err error) {
	match, err = utils.ComparePasswordAndHash(l.Password, hashPassword)
	if err != nil {
		return false, err
	}
	return match, err
}
