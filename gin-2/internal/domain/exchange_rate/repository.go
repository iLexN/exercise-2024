package exchange_rate

import "gorm.io/gorm"

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) GetAll() []ExchangeRate {

	var list []ExchangeRate

	r.Db.Find(&list)

	return list
}
