package env

import (
	"github.com/joho/godotenv"
	"log"
)

// JwtConfig gobal var, may use in defferent file
var JwtConfig JwtSetting

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	JwtConfig = CreateJwtSetting()
}
