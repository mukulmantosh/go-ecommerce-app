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

func (c Client) AddUserAddress(ctx context.Context, userAddress *models.UserAddress) (*models.UserAddress, error) {
	userAddress.ID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&userAddress)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &common_errors.ConflictError{}
		}
		if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
			return nil, &common_errors.ViolationError{}
		}
		return nil, result.Error
	}
	return userAddress, nil
}

func (c Client) GetUserAddressById(ctx context.Context, ID string) (*models.UserAddress, error) {
	userAddress := &models.UserAddress{}
	result := c.DB.WithContext(ctx).Where(&models.UserAddress{ID: ID}).First(&userAddress)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &common_errors.NotFoundError{}
		}
	}
	return userAddress, result.Error
}

func (c Client) UpdateUserAddress(ctx context.Context, userAddress *models.UserAddress) (bool, error) {
	result := c.DB.WithContext(ctx).Clauses(clause.Returning{}).
		Where(&models.UserAddress{ID: userAddress.ID}).Updates(
		&models.UserAddress{
			Address:    userAddress.Address,
			City:       userAddress.City,
			PostalCode: userAddress.PostalCode,
			Country:    userAddress.Country,
			Mobile:     userAddress.Mobile,
		})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return false, &common_errors.ConflictError{}
		}
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, &common_errors.NotFoundError{}
	}
	return true, nil
}

func (c Client) DeleteUserAddress(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(&models.UserAddress{ID: ID}).Error
}
