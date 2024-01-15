package env

import "testing"
import "github.com/joho/godotenv"

 func TestJwtSetting(t *testing.T){
	 
	 envFilePath := "../../.env"
	 // Load the custom .env file
	 err := godotenv.Load(envFilePath)
	 if err != nil {
		 t.Errorf("TestJwtSetting should not return an error, got: %v", err)
	 }
	 
	 
	jwtStruct := CreateJwtSetting()
	
	expectSecret := "my-secret-key111"
	expectIssuer := "test.local"
	
	if jwtStruct.Secret != expectSecret {
		t.Errorf("jwtStruct Secret mismatch. Expected: %s, got: %s", expectSecret, jwtStruct.Secret)
	}
	
	if jwtStruct.Issuer != expectIssuer {
		t.Errorf("jwtStruct Secret mismatch. Expected: %s, got: %s", expectIssuer, jwtStruct.Issuer)
	}
}