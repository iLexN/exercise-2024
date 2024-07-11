package user

import (
	"encoding/json"
	"gorm.io/datatypes"
	"payment-portal/internal/model"
	"payment-portal/internal/password"
)

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
