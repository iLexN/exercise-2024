package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go1/services/jwt_service"
	"go1/services/logger"
	"go1/usecase"
)

func LoginHandler(c *gin.Context) {
	var loginForm usecase.LoginForm

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Process the username and password
	// You can access the values using loginForm.Username and loginForm.Password

	msg := fmt.Sprintf("Received login request with username: %s, password: %s", loginForm.Username, loginForm.Password)
	logger.DefaultLogger.Info(msg)

	user := jwt_service.
		FromLoginForm(loginForm)

	jwtToken, err := jwt_service.
		CreateJwt(user)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":      "Login successful",
		"access_token": jwtToken,
	})
}
