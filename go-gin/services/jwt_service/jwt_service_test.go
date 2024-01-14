package jwt_service

import "testing"

func TestCreateJwt(t *testing.T) {
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

	if claims.Subject != userGenJwt.Username {
		t.Errorf("Role mismatch. Expected: %s, got: %s", claims.Subject, userGenJwt.Username)
	}

	if err != nil {
		t.Errorf("Decode should not return an error, got: %v", err)
	}
}
