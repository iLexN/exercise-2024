package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go1/services/logger"
	"time"
)

import "github.com/golang-jwt/jwt/v5"

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	var loginForm LoginForm

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Process the username and password
	// You can access the values using loginForm.Username and loginForm.Password

	msg := fmt.Sprintf("Received login request with username: %s, password: %s\n", loginForm.Username, loginForm.Password)
	logger.DefaultLogger.Info(msg)

	jwtToken := createJwt(loginForm.Username)

	c.JSON(200, gin.H{
		"message":      "Login successful",
		"access_token": jwtToken,
	})
}

type MyCustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func createJwt(username string) string {
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
			Subject:   username,
			//			ID:        "1",
			//			Audience:  []string{"somebody_else"},
		},
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Failed to sign the token:", err)
		return ""
	}

	fmt.Println("JWT Token:", signedToken)

	return signedToken
}
