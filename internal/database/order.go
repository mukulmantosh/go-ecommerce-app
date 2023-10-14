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

func (c Client) OrderPlacement(ctx context.Context, userId string) (*models.Order, error) {
	var cartDetails models.Cart
	c.DB.WithContext(ctx).Preload("Products").Where(&models.Cart{UserID: userId}).Find(&cartDetails)
	productCost := 0.0
	order := new(models.Order)
	order.ID = uuid.NewString()
	order.UserID = userId
	result := c.DB.WithContext(ctx).Create(&order)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &common_errors.ConflictError{}
		}
		return nil, result.Error
	}

	for _, product := range cartDetails.Products {
		productCost += product.Price
		c.DB.WithContext(ctx).Model(&order).Association("Products").Append(&models.Product{ID: product.ID})
	}

	c.DB.WithContext(ctx).Clauses(clause.Returning{}).Where(
		&models.Order{
			ID:          order.ID,
			UserID:      userId,
			Total:       float32(productCost),
			OrderStatus: "SUCCESS",
		})

	order.OrderStatus = "SUCCESS"
	order.Total = float32(productCost)
	c.DB.Save(&order)

	// Soft-delete cart
	c.DB.Delete(&models.Cart{}, &cartDetails.ID)
	return nil, nil
}
