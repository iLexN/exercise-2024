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

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		return
	}

}
