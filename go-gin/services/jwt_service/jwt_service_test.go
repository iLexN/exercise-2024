package jwt_service

import (
	"github.com/stretchr/testify/assert"
	"go1/services/env"
	"testing"
)

func TestCreateJwt(t *testing.T) {

	env.CreateJwtTestSetting()

	userGenJwt := UserGenJwt{
		Username: "user",
		Role:     "Admin",
	}

	token, err := CreateJwt(&userGenJwt)

	if token == "" {
		t.Error("CreateJwt should return a non-empty token")
	}
	if err != nil {
		t.Errorf("CreateJwt should not return an error, got: %v", err)
	}
}

func TestDecode(t *testing.T) {
	env.CreateJwtTestSetting()

	userGenJwt := UserGenJwt{
		Username: "user",
		Role:     "Admin",
	}

	// Create a test token
	tokenString, err := CreateJwt(&userGenJwt)

	if err != nil {
		t.Errorf("CreateJwt should not return an error, got: %v", err)
	}

	// Call the Decode function
	claims, err := Decode(tokenString)

	// Check if the claims are not nil and error is nil
	if claims == nil {
		t.Error("Decode should return non-nil claims")
	}

	if claims.Role != userGenJwt.Role {
		t.Errorf("Role mismatch. Expected: %s, got: %s", claims.Role, userGenJwt.Role)
	}

	assert.Equal(t, claims.Role, userGenJwt.Role, "jwt role should be same")

	if claims.Subject != userGenJwt.Username {
		t.Errorf("Role mismatch. Expected: %s, got: %s", claims.Subject, userGenJwt.Username)
	}

	assert.Equal(t, claims.Subject, userGenJwt.Username, "jwt sub shouldbe same")

	if err != nil {
		t.Errorf("Decode should not return an error, got: %v", err)
	}
}

func TestDecodeInvalidToken(t *testing.T) {
	env.CreateJwtTestSetting()

	// Create an invalid token string
	invalidTokenString := "invalid_token"

	// Call the Decode function with an invalid token
	claims, err := Decode(invalidTokenString)

	// Check if the claims are nil and the error is the expected invalid token error
	if claims != nil {
		t.Error("Decode should return nil claims for an invalid token")
	}
	if err == nil {
		t.Error("Decode should return error for an invalid token")
	}

	expectedErrorMsg := "token is malformed: token contains an invalid number of segments"
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)

}

func TestDecodeInvalidToken2(t *testing.T) {
	env.CreateJwtTestSetting()

	// Create an invalid token string
	invalidTokenString := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiQWRtaW4iLCJpc3MiOiJ0ZXN0LmxvY2FsIiwic3ViIjoiZ2lnaSIsImV4cCI6MTcwNTQ0Mjg5NiwibmJmIjoxNzA1MzU2NDk2LCJpYXQiOjE3MDUzNTY0OTZ9.AWIGyhdhfLv18NOjwJtrDOX0AA1WwF5w5lNPlFEnzJOXc_g1yfgkSYzfEspDVVrsfJSynt5Cov38PXdsa3CL-A"

	// Call the Decode function with an invalid token
	claims, err := Decode(invalidTokenString)

	// Check if the claims are nil and the error is the expected invalid token error
	if claims != nil {
		t.Error("Decode should return nil claims for an invalid token")
	}
	if err == nil {
		t.Error("Decode should return error for an invalid token")
	}

	expectedErrorMsg := "token signature is invalid: signature is invalid"
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)

}

func TestDecodeInvalidTokenExpired(t *testing.T) {
	env.CreateJwtTestSetting()

	// Create an invalid token string
	invalidTokenString := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiQWRtaW4iLCJpc3MiOiJ0ZXN0LmxvY2FsIiwic3ViIjoiZ2lnaSIsImV4cCI6MTcwNTI3MDU0OSwibmJmIjoxNzA1MzU2OTQ5LCJpYXQiOjE3MDUzNTY5NDl9.AukxKz1ZtmsijrXmpJj_40cl2tfNzBFBU_fKQBnngMyNmKDEJoRCzcA3d0dMeVmCqujawR62wHaafVuOxobe7A"

	// Call the Decode function with an invalid token
	claims, err := Decode(invalidTokenString)

	// Check if the claims are nil and the error is the expected invalid token error
	if claims != nil {
		t.Error("Decode should return nil claims for an invalid token")
	}
	if err == nil {
		t.Error("Decode should return error for an invalid token")
	}

	expectedErrorMsg := "token has invalid claims: token is expired"
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)

}
