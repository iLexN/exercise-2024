package usecase

import "testing"

func TestLoginFormStruct(t *testing.T) {
	form := LoginForm{
		Username: "user",
		Password: "pass",
	}

	if form.Username != "user" {
		t.Error("LoginForm username is not match")
	}

	if form.Password != "pass" {
		t.Error("LoginForm password is not match")
	}

	hash, err := form.HashPassword()

	if err != nil {
		t.Errorf("password hash should not return an error, got: %v", err)
	}

	if hash == "" {
		t.Error("password should return a non-empty string")
	}
}
