/*
	Copyright 2022 Google LLC

#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
*/
package abstract

import (
	"context"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
)

type Cart interface {
	NewCartForUser(ctx context.Context, cart *models.Cart) (*models.Cart, error)
	AddItemToCart(ctx context.Context, cartId string, productId string) (bool, error)
	GetCartInfoByUserID(ctx context.Context, userId string) (*models.Cart, int64, error)
}
