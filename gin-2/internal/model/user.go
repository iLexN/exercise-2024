package model

import (
	//    "gorm.io/gorm"
	"time"
)
import "gorm.io/datatypes"

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Roles     datatypes.JSON
}
