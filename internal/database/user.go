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
	"github.com/mukulmantosh/go-ecommerce-app/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
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

func (c Client) GetUserById(ctx context.Context, ID string) (*models.DisplayUser, error) {
	user := &models.User{}
	result := c.DB.WithContext(ctx).Preload(clause.Associations).Where(&models.User{ID: ID}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &common_errors.NotFoundError{}
		}
	}
	customUserDisplay := &models.DisplayUser{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return customUserDisplay, result.Error
}

func (c Client) UpdateUser(ctx context.Context, user *models.User) (bool, error) {
	params := &utils.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := utils.GenerateFromPassword(user.Password, params)
	if err != nil {
		log.Fatal(err)
	}

	result := c.DB.WithContext(ctx).Clauses(clause.Returning{}).
		Where(&models.User{ID: user.ID}).Updates(
		&models.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Password:  hash,
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

func (c Client) DeleteUser(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(&models.User{ID: ID}).Error
}
