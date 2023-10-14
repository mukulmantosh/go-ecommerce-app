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

func (c Client) OrderPlacement(ctx context.Context, userId string) (bool, error) {
	var cartDetails models.Cart
	var cartExist int64
	c.DB.WithContext(ctx).Preload("Products").Where(&models.Cart{UserID: userId}).Find(&cartDetails).Count(&cartExist)
	if cartExist == 0 {
		return false, &common_errors.CartEmptyError{}
	}
	productCost := 0.0
	order := new(models.Order)
	order.ID = uuid.NewString()
	order.UserID = userId
	result := c.DB.WithContext(ctx).Create(&order)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return false, &common_errors.ConflictError{}
		}
		return false, result.Error
	}

	for _, product := range cartDetails.Products {
		productCost += product.Price
		err := c.DB.WithContext(ctx).Model(&order).Association("Products").Append(&models.Product{ID: product.ID})
		if err != nil {
			return false, err
		}
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
	return true, nil
}
