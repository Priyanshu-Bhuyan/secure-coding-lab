package main

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString string) error {

	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		return fmt.Errorf("JWT_SECRET not configured")
	}

	_, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
	)

	return err
}
