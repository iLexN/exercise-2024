package main

import (
	"log"
	"log/slog"
	"payment-portal/internal/container"
	"payment-portal/internal/domain/gateway"
	"time"
)

func main() {
	di := container.NewContainer()

	// elsewhere in the application
	log.Printf("cfg: %#v", *di.Config.AppEnv) // Log the value of cfg
	slog.Info("hhiii", "di", *di)

	yesterday := time.Now().AddDate(0, 0, -1)
	list := di.GatewayRepository.GetAllWithEod(yesterday)

	exchangeRates := di.ExchangeRateRepository.GetAll()

	newII := gateway.CalGateways(list, exchangeRates)

	slog.Info("ggg", newII)

}
