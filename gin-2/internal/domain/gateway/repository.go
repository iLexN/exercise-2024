package gateway

import (
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) GetAllActive() []Gateway {
	var list []Gateway

	r.Db.Where("active = ?", true).
		Order("`order` ASC").
		Find(&list)

	return list
}

func (r *Repository) GetAllWithEod(date time.Time) []Gateway {

	var gateways []Gateway

	r.Db.Preload("Balances", func(db *gorm.DB) *gorm.DB {
		return db.Where("balance_at = ?", date.Format("2006-01-02"))
	}).Where("active = ?", true).Order("`order` ASC").Find(&gateways)

	return gateways
}
