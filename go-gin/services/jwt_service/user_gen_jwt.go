package jwt_service

import (
	"github.com/golang-jwt/jwt/v5"
	"go1/usecase"
)

type MyCustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type UserGenJwt struct {
	Username string
}

func FromLoginForm(loginForm usecase.LoginForm) *UserGenJwt {
	return &UserGenJwt{
		Username: loginForm.Username,
	}
}
