package token

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Parse(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // Check if the signing method is HMAC
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return nil, err
	}

	return claims, nil
}
