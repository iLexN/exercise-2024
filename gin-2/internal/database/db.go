package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"payment-portal/internal/config"
)

type Database struct {
	Db *gorm.DB
}

func NewConnection(c *config.Config) *Database {

	db, err := gorm.Open(mysql.Open(c.Database.Dns()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	return &Database{
		Db: db,
	}
}
