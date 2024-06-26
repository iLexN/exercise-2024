package password

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Hash(plainPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 12)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func Matches(plainPassword string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
