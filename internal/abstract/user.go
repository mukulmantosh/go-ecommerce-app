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

type User interface {
	AddUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserById(ctx context.Context, ID string) (*models.DisplayUser, error)
	UpdateUser(ctx context.Context, user *models.User) (bool, error)
	DeleteUser(ctx context.Context, ID string) error
}

type UserAddress interface {
	AddUserAddress(ctx context.Context, userAddress *models.UserAddress) (*models.UserAddress, error)
	GetUserAddressById(ctx context.Context, ID string) (*models.UserAddress, error)
	UpdateUserAddress(ctx context.Context, userAddress *models.UserAddress) (bool, error)
	DeleteUserAddress(ctx context.Context, ID string) error
}
