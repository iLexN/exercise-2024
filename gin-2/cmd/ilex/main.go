package main

import (
	"log"
	"log/slog"
	"payment-portal/internal/container"
	"time"
)

func main() {
	di := container.NewContainer()

	// elsewhere in the application
	log.Printf("cfg: %#v", *di.Config.AppEnv) // Log the value of cfg
	slog.Info("hhiii", "di", *di)

	yesterday := time.Now().AddDate(0, 0, -1)
	gateways := di.GatewayRepository.GetAllWithEod(yesterday)

	slog.Info("ggg", gateways)

}
