package abstract

import (
	"context"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
)

type User interface {
	AddUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserById(ctx context.Context, ID string) (*models.User, error)
}

type UserAddress interface {
	AddUserAddress(ctx context.Context, userAddress *models.UserAddress) (*models.UserAddress, error)
}
