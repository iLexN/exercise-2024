package jwt_service

import (
	"go1/usecase"
)

func FromLoginForm(loginForm usecase.LoginForm) *UserGenJwt {
	return &UserGenJwt{
		username: loginForm.Username,
	}
}
