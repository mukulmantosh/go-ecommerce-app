package database

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c Client) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.ID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &common_errors.ConflictError{}
		}
		return nil, result.Error
	}
	return user, nil
}

func (c Client) GetUserById(ctx context.Context, ID string) (*models.User, error) {
	user := &models.User{}
	result := c.DB.WithContext(ctx).Preload(clause.Associations).Where(&models.User{ID: ID}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &common_errors.NotFoundError{}
		}
	}
	return user, result.Error
}

//func (c Client) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
//	var products []models.Product
//	result := c.DB.WithContext(ctx).Clauses(clause.Returning{}).Where(
//		&models.Product{
//			Name:             product.Name,
//			Description:      product.Description,
//			Category:         product.Category,
//			ProductInventory: product.ProductInventory,
//		})
//
//	if result.Error != nil {
//		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
//			return nil, &common_errors.ConflictError{}
//		}
//		return nil, result.Error
//	}
//	if result.RowsAffected == 0 {
//		return nil, &common_errors.NotFoundError{}
//	}
//	return &products[0], nil
//}
//
//func (c Client) DeleteProduct(ctx context.Context, ID string) error {
//	return c.DB.WithContext(ctx).Delete(&models.Product{ProductID: ID}).Error
//}
