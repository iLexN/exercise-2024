package user

import (
	"gorm.io/gorm"
	"payment-portal/internal/model"
)

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) GetByEmailOrName(input string) (*model.User, error) {

	var u model.User

	result := r.Db.Where("email = ?", input).
		Or("name = ?", input).
		First(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}

func (r *Repository) GetById(id uint) (*model.User, error) {
	var u model.User

	result := r.Db.Where("id = ?", id).First(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}
