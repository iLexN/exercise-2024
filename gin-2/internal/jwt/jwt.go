package jwt

import (
	"payment-portal/internal/config"
	"time"
)

type TokenInfo struct {
	SignedToken string
	ExpireAt    time.Time
}

type TokenServices struct {
	Config *config.JwtConfig
}

func (t *TokenServices) CreateToken(u int) (TokenInfo, error) {
	// Define
}
