package container

import (
	"log/slog"
	"os"
	"payment-portal/internal/config"
	"payment-portal/internal/database"
)

type Container struct {
	Config *config.Config
	Logger *slog.Logger
	Db *database.Database
}

func NewContainer() *Container {
	cfg := config.NewConfig()
	
	return &Container{
		Config: cfg,
		Logger: newLogger(),
		Db:     database.NewConnection(cfg),
	}
}

func newLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	return logger
}
