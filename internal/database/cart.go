package database

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"gorm.io/gorm"
)

func (c Client) NewCartForUser(ctx context.Context, cart *models.Cart) (*models.Cart, error) {
	cart.ID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&cart)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &common_errors.ConflictError{}
		}
		return nil, result.Error
	}
	return cart, nil
}

func (c Client) AddItemToCart(ctx context.Context, cartId string, productId string) (bool, error) {
	var cartInfo models.Cart
	var productInfo models.Product
	c.DB.WithContext(ctx).Where(&models.Cart{ID: cartId}).First(&cartInfo)
	c.DB.WithContext(ctx).Where(&models.Product{ID: productId}).First(&productInfo)

	err := c.DB.WithContext(ctx).Model(&cartInfo).Association("Products").Append(&models.Product{ID: productInfo.ID})
	if err != nil {
		return false, err
	}

	return true, err
}
