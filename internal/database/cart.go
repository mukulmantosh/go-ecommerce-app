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

func (c Client) GetCartInfoByUserID(ctx context.Context, userId string) (*models.Cart, int64, error) {
	var cartInfo models.Cart
	var cartPresent int64
	c.DB.WithContext(ctx).Where(&models.Cart{UserID: userId}).First(&cartInfo).Count(&cartPresent)
	return &cartInfo, cartPresent, nil

}
