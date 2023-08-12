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

func (c Client) UpdateUser(ctx context.Context, user *models.User) (bool, error) {
	result := c.DB.WithContext(ctx).Clauses(clause.Returning{}).
		Where(&models.User{ID: user.ID}).Updates(
		&models.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Password:  user.Password,
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
