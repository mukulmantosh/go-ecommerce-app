package abstract

import (
	"context"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
)

type Category interface {
	AddCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	GetCategoryById(ctx context.Context, ID string) (*models.Category, error)
	UpdateCategory(ctx context.Context, category *models.Category) (bool, error)
	DeleteCategory(ctx context.Context, ID string) error
}
