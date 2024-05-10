package config

import (
	"payment-portal/internal/env"
	"time"
)

type serverConfig struct {
	Port            int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

func newServerConfig() *serverConfig {

	readTimeout := time.Duration(env.GetInt("SERVER_READ_TO", 5)) * time.Second
	writeTimeout := time.Duration(env.GetInt("SERVER_WRITE_TO", 10)) * time.Second
	shutdownTimeout := time.Duration(env.GetInt("SERVER_SHUTDOWN_TO", 5)) * time.Second

	return &serverConfig{
		Port:            env.GetInt("SERVER_PORT", 9000),
		ReadTimeout:     readTimeout,
		WriteTimeout:    writeTimeout,
		ShutdownTimeout: shutdownTimeout,
	}
}
