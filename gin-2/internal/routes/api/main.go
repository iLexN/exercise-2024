package main

import (
	"log"
	"log/slog"
	"payment-portal/internal/container"
	"payment-portal/internal/server"
)

func main() {
	di := container.NewContainer()

	// elsewhere in the application
	log.Printf("cfg: %#v", *di.Config.AppEnv) // Log the value of cfg
	slog.Info("hhiii", "di", *di)

	server.Run(di)
}
