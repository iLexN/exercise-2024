package container

import (
	"log/slog"
	"os"
	"payment-portal/internal/config"
	"payment-portal/internal/database"
	"payment-portal/internal/domain/exchange_rate"
	"payment-portal/internal/domain/gateway"
	"payment-portal/internal/domain/transaction"
	"payment-portal/internal/domain/uiprocess"
	"payment-portal/internal/domain/user"
	"payment-portal/internal/jwt"
)

type Container struct {
	Config                 *config.Config
	Logger                 *slog.Logger
	UserRepository         *user.Repository
	TransactionRepository  *transaction.Repository
	ExchangeRateRepository *exchange_rate.Repository
	UipathRepository       *uiprocess.Repository
	GatewayRepository      *gateway.Repository
	JwtTokenServices       *jwt.TokenServices
}

func NewContainer() *Container {
	cfg := config.NewConfig()
	db := database.NewConnection(cfg)

	// repository
	userRepository := user.Repository{Db: db.Db}
	transactionRepository := transaction.Repository{Db: db.Db}
	exchangeRateRepository := exchange_rate.Repository{Db: db.Db}
	gatewayRepository := gateway.Repository{Db: db.Db}
	uiprocessRepository := uiprocess.Repository{Db: db.Db}

	logger := newLogger()
	jwtTokenService := jwt.NewTokenServices(cfg.JwtConfig)

	return &Container{
		Config:                 cfg,
		Logger:                 logger,
		UserRepository:         &userRepository,
		TransactionRepository:  &transactionRepository,
		ExchangeRateRepository: &exchangeRateRepository,
		UipathRepository:       &uiprocessRepository,
		GatewayRepository:      &gatewayRepository,
		JwtTokenServices:       jwtTokenService,
	}
}

func newLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	return logger
}
