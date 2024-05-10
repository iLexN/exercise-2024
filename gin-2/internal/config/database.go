package config

import (
	"fmt"
	"payment-portal/internal/env"
)

type databaseConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func (mysql *databaseConfig) Dns() string {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysql.Username,
		mysql.Password,
		mysql.Host,
		mysql.Port,
		mysql.Database,
	)
}

func NewMysql() *databaseConfig {
	return &databaseConfig{
		Host:     env.GetString("DB_HOST", "127.0.0.1"),
		Port:     env.GetString("DB_PORT", "3306"),
		Database: env.GetString("DB_DATABASE", "payment-portal"),
		Username: env.GetString("DB_USERNAME", "payment-portal"),
		Password: env.GetString("DB_PASSWORD", "payment-portal"),
	}
}
