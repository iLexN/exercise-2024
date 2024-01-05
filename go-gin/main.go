package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go1/services/logger"
	"go1/services/route_service"
)

//var DefaultLogger logger.Logger = logger.CreateZeroLog()

func main() {

	logger.DefaultLogger = logger.CreateZeroLog()

	// Create a new Gin router
	router := gin.Default()

	// set trusted proxies , nil to  disable
	trustedProxies := []string{"127.0.0.1"}
	err := router.SetTrustedProxies(trustedProxies)
	if err != nil {
		fmt.Print(err)
		return
	}

	route_service.CreateRoute(router)

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		return
	}

}
