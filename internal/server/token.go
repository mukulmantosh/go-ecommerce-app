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
