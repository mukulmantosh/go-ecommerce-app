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
package database

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"os"
	"time"
)

func (c Client) Login(ctx context.Context, user *models.Login) (string, error) {
	var userCount int64
	userInfo := &models.User{}

	c.DB.WithContext(ctx).Where(&models.User{Username: user.Username}).First(&userInfo).Count(&userCount)
	if userCount > 0 {
		passwordMatch, err := user.VerifyPassword(userInfo.Password)
		if err != nil {
			return "", &common_errors.NotFoundError{}
		}
		if !passwordMatch {
			return "", &common_errors.NotFoundError{}
		} else {
			// return JWT token

			// Set custom claims with expiry of 30 minutes
			claims := &models.CustomJWTClaims{
				Name:     userInfo.FirstName + " " + userInfo.LastName,
				UserName: userInfo.Username,
				UserID:   userInfo.ID,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
				},
			}

			// Create token with claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			// Generate encoded token and send it as response.
			jwtSecret := os.Getenv("JWT_SECRET")
			tokenInfo, err := token.SignedString([]byte(jwtSecret))
			if err != nil {
				return "", err
			}
			return tokenInfo, nil

		}
	} else {
		return "", &common_errors.NotFoundError{}
	}
}
