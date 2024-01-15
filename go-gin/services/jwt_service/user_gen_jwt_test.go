package jwt_service

import (
	"go1/usecase"
	"testing"
)

func TestFromLoginForm(t *testing.T) {
	loginForm := usecase.LoginForm{
		Username: "username",
		Password: "password",
	}

	expected := &UserGenJwt{
		Username: "username",
		Role:     "Admin",
	}

	result := FromLoginForm(loginForm)

	if result.Username != expected.Username {
		t.Errorf("Username mismatch. Expected: %s, got: %s", expected.Username, loginForm.Username)
	}

	if result.Role != expected.Role {
		t.Errorf("Role mismatch. Expected: %s, got: %s", expected.Username, loginForm.Username)
	}

}
