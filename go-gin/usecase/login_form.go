package usecase

import "go1/services/password"

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (f *LoginForm) HashPassword() (string, error) {
	hashedPassword, err := password.Hash(f.Username)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}
