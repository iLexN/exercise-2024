package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"payment-portal/internal/container"
	"payment-portal/internal/routes"
	"syscall"
)

func Run(container *container.Container) {

	port := container.Config.Server.Port
	readTimeout := container.Config.Server.ReadTimeout
	writeTimeout := container.Config.Server.WriteTimeout

	router := gin.Default()

	routes.Setup(router, container)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Handler:      router,
	}

	// Send log message.
	servereServe(server, container)
}

func servereServe(server *http.Server, container *container.Container) {

	port := container.Config.Server.Port
	shutdownTimeout := container.Config.Server.ShutdownTimeout

	slog.Info("Starting server...", "port", port)
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	//	select {
	//	case <-ctx.Done():
	//		log.Printf("timeout of %.2f seconds.", shutdownTimeout.Seconds())
	//	}
	<-ctx.Done()
	log.Printf("timeout of %.2f seconds.", shutdownTimeout.Seconds())
	log.Println("Server exiting")
}
