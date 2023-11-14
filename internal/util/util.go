package util

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId uint) (*string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Subject:   strconv.Itoa(int(userId)),
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

func ConvertObjToStringReader[T any](obj T) io.Reader {
	result, _ := json.Marshal(&obj)
	return strings.NewReader(string(result))
}
