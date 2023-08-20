package abstract

import (
	"context"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
)

type Cart interface {
	NewCartForUser(ctx context.Context, cart *models.Cart) (*models.Cart, error)
	AddItemToCart(ctx context.Context, cartId string, productId string) (bool, error)
}
