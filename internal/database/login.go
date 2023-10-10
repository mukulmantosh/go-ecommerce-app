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
	fmt.Println(userCount)
	if userCount > 0 {
		fmt.Println(userInfo.FirstName, userInfo.LastName, userInfo.Password)
		passwordMatch, err := user.VerifyPassword(userInfo.Password)
		fmt.Println(passwordMatch)
		if err != nil {
			fmt.Println("here 1")
			return "", &common_errors.NotFoundError{}
		}
		if passwordMatch != true {
			fmt.Println("here 2")
			return "", &common_errors.NotFoundError{}
		} else {
			// return JWT token

			// Set custom claims with expiry of 3 hours
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
			tokenInfo, err := token.SignedString([]byte("secret"))
			if err != nil {
				fmt.Println("here 3")
				return "", err
			}
			return tokenInfo, nil

		}
	} else {
		fmt.Println("came here last")
		return "", &common_errors.NotFoundError{}
	}
}
