package model

import "gorm.io/gorm"
import "gorm.io/datatypes"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Roles    datatypes.JSON
}
