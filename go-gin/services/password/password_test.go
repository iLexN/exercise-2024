package password

import "testing"

const password_string = "this-is-password"

func TestHash(t *testing.T) {
	password := password_string

	result, err := Hash(password)

	if err != nil {
		t.Errorf("password hash should not return an error, got: %v", err)
	}

	if result == "" {
		t.Error("password should return a non-empty string")
	}

}

func TestGolangHashMatches(t *testing.T) {
	hashString := "$2a$12$lt2kAQW5K.oqNm.vIpv3g.PEigrrgaod1j3figuZRDRHa.7Xr2QwC"

	plainPassword := password_string

	result, err := Matches(plainPassword, hashString)

	if err != nil {
		t.Errorf("password match should not return an error, got: %v", err)
	}

	if result != true {
		t.Error("password should return a true")
	}

}

func TestPhpHashMatches(t *testing.T) {
	hashString := "$2y$10$aCA3iwDtfpDd3a6FSSzFcuOuPbXSjzMBBUp7vBRmDAZjynGCOWQXW"

	plainPassword := password_string

	result, err := Matches(plainPassword, hashString)

	if err != nil {
		t.Errorf("password match should not return an error, got: %v", err)
	}

	if result != true {
		t.Error("password should return a true")
	}

}
