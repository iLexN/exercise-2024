package model

import (
	//    "gorm.io/gorm"
	"time"
)
import "gorm.io/datatypes"

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	Roles     datatypes.JSON
}
