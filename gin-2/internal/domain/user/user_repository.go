package user

import (
	"gorm.io/gorm"
	"payment-portal/internal/model"
)

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) GetByEmailOrName(input string) (*model.User, error) {

	u := model.User{}

	result := r.Db.Where("email = ?", input).
		Or("name = ?", input).
		First(&u)

	return &u, result.Error
}
