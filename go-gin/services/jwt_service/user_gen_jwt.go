package jwt_service

import (
	"go1/usecase"
)

type UserGenJwt struct {
	Username string
}

func FromLoginForm(loginForm usecase.LoginForm) *UserGenJwt {
	return &UserGenJwt{
		Username: loginForm.Username,
	}
}
