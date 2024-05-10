package config

import "payment-portal/internal/env"

type appEnv struct {
	Env string
}

func newAppEnv() *appEnv {
	return &appEnv{
		Env: env.GetString("APP_ENV", "production"),
	}
}
