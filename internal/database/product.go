package database

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-ecommerce-app/internal/generic/common_errors"
	"go-ecommerce-app/internal/models"
	"gorm.io/gorm"
)

func (c Client) AllProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).Find(&products)
	return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &common_errors.ConflictError{}
		}
		return nil, result.Error
	}
	return product, nil
}
