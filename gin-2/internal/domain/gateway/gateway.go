package gateway

import (
	"gorm.io/datatypes"
	"payment-portal/internal/domain/Balance"
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

	Balances []Balance.Balance `gorm:"foreignKey:GatewaysID"`
}

type Summary struct {
	Gateway
	//	Balance     map[string]map[string]interface{}
	LastUpdated time.Time
}

type CalResult struct {
	CalAllBalance float64
	Currency      []map[string]interface{}
	Gateways      []Summary
}

func (g *Gateway) ToDisplay() map[string]interface{} {
	return map[string]interface{}{
		"id":           g.ID,
		"name":         g.Name,
		"display_name": g.DisplayName,
	}
}
