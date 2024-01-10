package route_service

import (
	"github.com/gin-gonic/gin"
	controller "go1/controller"
	"go1/middleware"
	"net/http"
)

func CreateRoute(r *gin.Engine) {
	// Define a route and its handler function
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.POST("/login", controller.LoginHandler)
	r.GET("/try-check-jwt", middleware.AuthJwt(), controller.CheckJwt)

	r.GET("/grpc-hi", controller.HelloGrpc)
	r.GET("/grpc-hi-php", controller.HelloPhpGrpc)

	r.GET("/hi-php-api", controller.PhpApiCall)
}
