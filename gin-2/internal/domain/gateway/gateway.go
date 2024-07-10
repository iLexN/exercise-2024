package gateway

import (
	"gorm.io/datatypes"
	"time"
)

type Gateway struct {
	ID           uint `gorm:"primarykey"`
	Name         string
	DisplayName  string
	Config       datatypes.JSON
	ClientConfig datatypes.JSON
	active       bool
	syncMethod   int
	Order        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (g *Gateway) ToDisplay() map[string]interface{} {
	return map[string]interface{}{
		"id":           g.ID,
		"name":         g.Name,
		"display_name": g.DisplayName,
	}
}
