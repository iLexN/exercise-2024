package route_service

import (
	"github.com/gin-gonic/gin"
	controllers "go1/controller"
	"net/http"
)

func CreateRoute(r *gin.Engine) {
	// Define a route and its handler function
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.POST("/login", controllers.LoginHandler)
}

