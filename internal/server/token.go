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
package server

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"os"
)

func ParseJWTToken(tokenString string) (*models.CustomJWTClaims, error) {
	// Parse the JWT token with custom claims
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Replace this with your own secret key or public key (if using asymmetric signing)
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.CustomJWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
