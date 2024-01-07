package env

import (
	"github.com/joho/godotenv"
	"log"
)

var JwtConfig JwtSetting

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	JwtConfig = CreateJwtSetting()
}
