package abstract

import (
	"context"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
)

type Order interface {
	OrderPlacement(ctx context.Context, userId string) (bool, error)
	OrderList(ctx context.Context, userId string) ([]models.Order, error)
}
