package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken() (*string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte("MySignature"))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func ValidateToken(token string) error {

	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("MySignature"), nil
	})

	return err
}