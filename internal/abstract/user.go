package abstract

import (
	"context"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
)

type User interface {
	AddUser(ctx context.Context, user *models.User) (*models.User, error)
	AddUserAddress(ctx context.Context, userAddress *models.UserAddress) (*models.UserAddress, error)
}
