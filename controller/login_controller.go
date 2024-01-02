package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go1/services/logger"
)

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
	fmt.Printf("Received login request with username: %s, password: %s\n", loginForm.Username, loginForm.Password)

	logger.DefaultLogger.Info("hello")

	c.JSON(200, gin.H{"message": "Login successful"})
}
