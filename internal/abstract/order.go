package abstract

import (
	"context"
)

type Order interface {
	OrderPlacement(ctx context.Context, userId string) (bool, error)
}
