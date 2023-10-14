package database

import (
	"fmt"
	"github.com/mukulmantosh/go-ecommerce-app/internal/abstract"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		"localhost", "postgres", "mukul123", "ecommerce", 5432, "disable")

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
