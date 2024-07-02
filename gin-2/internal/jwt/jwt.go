package jwt

import (
	"errors"
	"payment-portal/internal/config"
	"payment-portal/internal/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var errInvalidToken = errors.New("invalid token")

type TokenInfo struct {
	SignedToken string
	ExpireAt    time.Time
}

type TokenServices struct {
	Config *config.JwtConfig
}

func NewTokenServices(config *config.JwtConfig) *TokenServices {
	return &TokenServices{
		Config: config,
	}
}

func (t *TokenServices) getSecretKey() []byte {
	return []byte(t.Config.Secret)
}
func (t *TokenServices) getIssuer() string {
	return t.Config.Issuer
}
func (t *TokenServices) getExpireHour() time.Duration {
	return time.Duration(t.Config.ExpireHour)
}

func (t *TokenServices) CreateToken(u *model.User) (*TokenInfo, error) {

	expire := time.Now().Add(t.getExpireHour() * time.Hour)

	claims := jwt.RegisteredClaims{
		// A usual scenario is to set the expiration time relative to the current time
		ExpiresAt: jwt.NewNumericDate(expire),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    t.getIssuer(),
		Subject:   u.Name,
		ID:        strconv.Itoa(int(u.ID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(t.getSecretKey())

	if err != nil {
		return &TokenInfo{}, err
	}

	return &TokenInfo{
		SignedToken: signedToken,
		ExpireAt:    expire,
	}, nil
}

func (t *TokenServices) DecodeToken(tokenString string) (*jwt.RegisteredClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Provide the key or public key to validate the token's signature
		return t.getSecretKey(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errInvalidToken
}
