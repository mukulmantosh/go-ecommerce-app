package database

import (
	"context"
	"fmt"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient interface {
	Ready() bool
	RunMigration() error
	AllProducts(ctx context.Context) ([]models.Product, error)
	AddProduct(ctx context.Context, product *models.Product) (*models.Product, error)
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
		&models.Product{},
		&models.Category{},
		&models.ProductInventory{},
		&models.User{},
		&models.UserAddress{},
		&models.OrderSession{},
	)
	if err != nil {
		return err
	}
	return nil
}
