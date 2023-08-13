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

func (c Client) AddCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	category.ID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &common_errors.ConflictError{}
		}
		if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
			return nil, &common_errors.ViolationError{}
		}
		return nil, result.Error
	}
	return category, nil
}

func (c Client) GetCategoryById(ctx context.Context, ID string) (*models.Category, error) {
	category := &models.Category{}
	result := c.DB.WithContext(ctx).Where(&models.Category{ID: ID}).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &common_errors.NotFoundError{}
		}
	}
	return category, result.Error
}

func (c Client) UpdateCategory(ctx context.Context, category *models.Category) (bool, error) {
	result := c.DB.WithContext(ctx).Clauses(clause.Returning{}).
		Where(&models.Category{ID: category.ID}).Updates(
		&models.Category{
			Name:        category.Name,
			Description: category.Description,
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

func (c Client) DeleteCategory(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(&models.Category{ID: ID}).Error
}
