package container

import (
	"log/slog"
	"os"
	"payment-portal/internal/config"
	"payment-portal/internal/database"
	"payment-portal/internal/domain/user"
	"payment-portal/internal/jwt"
)

type Container struct {
	Config           *config.Config
	Logger           *slog.Logger
	UserRepository   *user.Repository
	JwtTokenServices *jwt.TokenServices
}

func NewContainer() *Container {
	cfg := config.NewConfig()
	db := database.NewConnection(cfg)
	userRepository := user.Repository{Db: db.Db}

	return &Container{
		Config:           cfg,
		Logger:           newLogger(),
		UserRepository:   &userRepository,
		JwtTokenServices: jwt.NewTokenServices(cfg.JwtConfig),
	}
}

func newLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	return logger
}
