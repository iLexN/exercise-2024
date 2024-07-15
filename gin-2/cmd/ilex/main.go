package main

import (
	"log"
	"log/slog"
	"payment-portal/internal/container"
)

func main() {
	di := container.NewContainer()

	// elsewhere in the application
	log.Printf("cfg: %#v", *di.Config.AppEnv) // Log the value of cfg
	slog.Info("hhiii", "di", *di)

	ui, err := di.UserRepository.GetById(100)

	if err != nil {
		slog.Info("eeeee", "eeeee", err.Error())
		return
	}

	slog.Info(ui.Name)

}
