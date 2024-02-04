package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

func open(config *DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", config.dsn())
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	return db, nil
}

type DatabaseConfig struct {
	Username     string
	Password     string
	Host         string
	DatabaseName string
}

func (c *DatabaseConfig) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.Username, c.Password, c.Host, c.DatabaseName)
}

func CreateFromEnv() *DatabaseConfig {
	return &DatabaseConfig{
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASS"),
		Host:         os.Getenv("DB_HOST"),
		DatabaseName: os.Getenv("DB_NAME"),
	}
}

func (c *DatabaseConfig) Open() (*sqlx.DB, error) {
	db, err := open(c)
	if err != nil {
		return nil, err
	}
	return db, err
}
