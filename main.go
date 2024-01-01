package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// set trusted proxies , nil to  disable
	trustedProxies := []string{"127.0.0.1"}
	err := router.SetTrustedProxies(trustedProxies)
	if err != nil {
		fmt.Print(err)
		return
	}

	// Define a route and its handler function
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	router.POST("/login", func(c *gin.Context) {
		var loginInfo LoginInfo

		if err := c.ShouldBindJSON(&loginInfo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
		}

		fmt.Printf(
			"Got the login with %s (%s)\n",
			loginInfo.Username,
			loginInfo.Password,
		)

		c.JSON(200, gin.H{"message": "OK"})
	})

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		return
	}

}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
