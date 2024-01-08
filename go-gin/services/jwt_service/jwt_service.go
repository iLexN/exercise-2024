package jwt_service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go1/services/env"
	"go1/services/logger"
	"time"
)

func CreateJwt(u *UserGenJwt) (string, error) {
	// Define the secret key used to sign the token
	secretKey := []byte(env.JwtConfig.Secret)

	// Define the payload of the token
	claims := MyCustomClaims{
		"Admin",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    env.JwtConfig.Issuer,
			Subject:   u.Username,
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

func Decode(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Provide the key or public key to validate the token's signature
		return []byte(env.JwtConfig.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

type MyCustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}
