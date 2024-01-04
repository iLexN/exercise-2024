package jwt_service

import (
	"github.com/golang-jwt/jwt/v5"
	"go1/services/logger"
	"time"
)

type MyCustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type UserGenJwt struct {
	username string
}

func (g UserGenJwt) CreateJwt() (string, error) {
	// Define the secret key used to sign the token
	secretKey := []byte("my-secret-key111")

	// Define the payload of the token
	claims := MyCustomClaims{
		"Admin",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test.local",
			Subject:   g.username,
			//			ID:        "1",
			//			Audience:  []string{"somebody_else"},
		},
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		logger.DefaultLogger.Error("Failed to sign the token:", err)
		return "", err
	}

	return signedToken, nil
}
