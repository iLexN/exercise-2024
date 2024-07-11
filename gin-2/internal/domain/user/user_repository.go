package user

import (
	"encoding/json"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"payment-portal/internal/model"
	"payment-portal/internal/password"
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

func (r *Repository) CreateUser(input *CreateUserInput) (*model.User, error) {

	hashPassword, err := password.Hash(input.Password)

	if err != nil {
		return nil, err
	}

	roleJSON, err := json.Marshal([]Roles{input.Role})
	if err != nil {
		return nil, err
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashPassword,
		Roles:    datatypes.JSON(roleJSON),
	}

	result := r.Db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}
