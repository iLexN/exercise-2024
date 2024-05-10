package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	gowebly "github.com/gowebly/helpers"
)

// runServer runs a new HTTP server with the loaded environment variables.
func runServer() error {
	// Validate environment variables.
	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "9000"))
	if err != nil {
		return err
	}

	// Create a new Fiber server.
	router := gin.Default()

	// Handle static files.
	router.Static("/static", "./static")

	// Handle index page view.
	router.GET("/", indexViewHandler)

	// Handle API endpoints.
	router.GET("/api/hello-world", showContentAPIHandler)

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	// Send log message.
	slog.Info("Starting server...", "port", port)

	return server.ListenAndServe()
}
