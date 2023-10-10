package abstract

import (
	"context"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
)

type Authenticate interface {
	Login(ctx context.Context, user *models.Login) (string, error)
}
