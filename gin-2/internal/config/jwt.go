package config

import "payment-portal/internal/env"

type JwtConfig struct {
	Secret string
	Issuer string
}

func newJwtConfig() *JwtConfig {
	return &JwtConfig{
		Secret:  env.GetString("JWT_SECRET", "30624700"),
		Issuer:  env.GetString("JWT_ISSUER", "iLexN"),
	}
}