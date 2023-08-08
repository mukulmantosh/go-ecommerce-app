package abstract

import (
	"context"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
)

type Product interface {
	AllProducts(ctx context.Context) ([]models.Product, error)
	AddProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	GetProductById(ctx context.Context, ID string) (*models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	DeleteProduct(ctx context.Context, ID string) error
}
