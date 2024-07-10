package gateway

import "gorm.io/gorm"

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
