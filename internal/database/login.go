package database

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mukulmantosh/go-ecommerce-app/internal/generic/common_errors"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"time"
)

func (c Client) Login(ctx context.Context, user *models.Login) (string, error) {
	var userCount int64
	userInfo := &models.User{}

	c.DB.WithContext(ctx).Where(&models.User{Username: user.Username}).First(&userInfo).Count(&userCount)
	if userCount > 0 {
		fmt.Println(userInfo.FirstName, userInfo.LastName, userInfo.Password)
		passwordMatch, err := user.VerifyPassword(userInfo.Password)
		if err != nil {
			return "", &common_errors.NotFoundError{}
		}
		if passwordMatch != true {
			return "", &common_errors.NotFoundError{}
		} else {
			// return JWT token

			// Set custom claims with expiry of 3 hours
			claims := &models.CustomJWTClaims{
				Name: user.Username,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
				},
			}

			// Create token with claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			// Generate encoded token and send it as response.
			tokenInfo, err := token.SignedString([]byte("secret"))
			if err != nil {
				return "", err
			}
			return tokenInfo, nil

		}
	} else {
		return "", &common_errors.NotFoundError{}
	}
}
