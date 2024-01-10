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
package database

import (
	"fmt"
	"github.com/mukulmantosh/go-ecommerce-app/internal/abstract"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

type DBClient interface {
	Ready() bool
	RunMigration() error
	CloseConnection()
	abstract.Authenticate
	abstract.User
	abstract.UserAddress
	abstract.Product
	abstract.Category
	abstract.Cart
	abstract.Order
}

type Client struct {
	DB *gorm.DB
}

func NewDBClient() (DBClient, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	databasePort, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal("Invalid DB Port")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbHost, dbUsername, dbPassword, dbName, databasePort, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	client := Client{DB: db}
	return client, nil
}

func NewTestDBClient() (DBClient, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	client := Client{DB: db}
	return client, nil

}

func (c Client) Ready() bool {

	var ready string
	result := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if result.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}

func (c Client) RunMigration() error {
	err := c.DB.AutoMigrate(
		&models.User{},
		&models.UserAddress{},
		&models.Cart{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) CloseConnection() {
	sqlDB, _ := c.DB.DB()
	err := sqlDB.Close()
	if err != nil {
		return
	}
}
