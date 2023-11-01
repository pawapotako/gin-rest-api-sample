package util

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(UserId uuid.UUID) (*string, error) {

	con := loadConfig()
	privateBytes := []byte(con.Cryption.PrivateKey)

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Create a new token object, specifying the signing method and claims
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set the desired claims for the token
	claims["UserId"] = UserId
	claims["ExpiresAt"] = time.Now().Add(time.Hour * 1).Unix()

	// Generate the token with the private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func GenerateRefreshToken(UserId uuid.UUID) (*string, error) {

	con := loadConfig()
	privateBytes := []byte(con.Cryption.PrivateKey)

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Create a new token object, specifying the signing method and claims
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set the desired claims for the token
	claims["UserId"] = UserId
	claims["ExpiresAt"] = time.Now().AddDate(0, 1, 0).Unix()

	// Generate the token with the private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func ValidateToken(tokenString string) (*uuid.UUID, *int64, error) {

	con := loadConfig()
	publicBytes := []byte(con.Cryption.PublicKey)

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	// Parse the token with the public key to get the claims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token's signing method is RSA (RS256)
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Check if the token is valid and has valid claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract UserId and ExpiresAt from claims
		userIdStr, userIdExists := claims["UserId"].(string)
		expiresAtFloat64, expiresAtExists := claims["ExpiresAt"].(float64)

		if !userIdExists || !expiresAtExists {
			return nil, nil, fmt.Errorf("token claims are not valid")
		}

		userId, err := uuid.FromString(userIdStr)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse UserId from token claims: %w", err)
		}

		expiresAt := int64(expiresAtFloat64)

		if expiresAt < time.Now().Unix() {
			return nil, nil, fmt.Errorf("token is expired")
		}

		return &userId, &expiresAt, nil

	}

	return nil, nil, fmt.Errorf("token expired")
}
