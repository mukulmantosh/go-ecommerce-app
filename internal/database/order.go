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

func (c Client) OrderList(ctx context.Context, userId string) ([]models.Order, error) {
	var orders []models.Order
	result := c.DB.WithContext(ctx).Preload("Products").Where(&models.Order{UserID: userId}).Find(&orders)
	return orders, result.Error
}
