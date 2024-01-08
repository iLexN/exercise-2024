package env

import (
	"os"
)

type JwtSetting struct {
	Secret string
	Issuer string
}

func CreateJwtSetting() JwtSetting {
	return JwtSetting{
		Secret: os.Getenv("JWT_SECRET"),
		Issuer: os.Getenv("JWT_ISSUER"),
	}
}
