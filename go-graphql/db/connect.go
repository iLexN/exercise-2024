package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
	"time"
)

func Open(config *DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", config.dsn())
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)

	return db, nil
}

type DatabaseConfig struct {
	Username        string
	Password        string
	Host            string
	DatabaseName    string
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
}

func (c *DatabaseConfig) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.Username, c.Password, c.Host, c.DatabaseName)
}

func CreateFromEnv() *DatabaseConfig {
	connMaxLifetimeSeconds, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME"))
	maxIdleConns, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	maxOpenConns, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))

	connMaxLifetime := time.Second * time.Duration(connMaxLifetimeSeconds)

	return &DatabaseConfig{
		Username:        os.Getenv("DB_USER"),
		Password:        os.Getenv("DB_PASS"),
		Host:            os.Getenv("DB_HOST"),
		DatabaseName:    os.Getenv("DB_NAME"),
		ConnMaxLifetime: time.Second * connMaxLifetime,
		MaxIdleConns:    maxIdleConns,
		MaxOpenConns:    maxOpenConns,
	}
}

func (c *DatabaseConfig) Open() (*sqlx.DB, error) {
	db, err := Open(c)
	if err != nil {
		return nil, err
	}
	return db, err
}
