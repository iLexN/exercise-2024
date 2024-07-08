package main

import (
	"fmt"
	"log"
	"log/slog"
	"payment-portal/internal/container"
)

func main() {
	di := container.NewContainer()

	// elsewhere in the application
	log.Printf("cfg: %#v", *di.Config.AppEnv) // Log the value of cfg
	slog.Info("hhiii", "di", *di)

	exchangeList := di.ExchangeRateRepository.GetAll()

	for _, exchange := range exchangeList {
		fmt.Printf("ID: %d, From: %s, To: %s, Rate: %.6f\n",
			exchange.ID, exchange.FromCurrency, exchange.ToCurrency, exchange.Rate)
	}
}
