package container

import (
	"log/slog"
	"os"
	"payment-portal/internal/config"
	"payment-portal/internal/database"
	"payment-portal/internal/domain/user"
)

type Container struct {
	Config         *config.Config
	Logger         *slog.Logger
//	Db             *database.Database
	UserRepository *user.Repository
}

func NewContainer() *Container {
	cfg := config.NewConfig()
	db := database.NewConnection(cfg)
	userRepository := user.Repository{Db: db.Db}

	return &Container{
		Config:         cfg,
		Logger:         newLogger(),
//		Db:             db,
		UserRepository: &userRepository,
	}
}

func newLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	return logger
}
