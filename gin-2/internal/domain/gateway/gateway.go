package gateway

import (
	"gorm.io/datatypes"
	"math"
	"payment-portal/internal/domain/Balance"
	"sort"
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
	Currency      []struct {
		Currency  string
		CalAmount float64
	}
	Gateways []Summary
}

func (r *CalResult) CurrencyToDisplay() []struct {
	Currency  string
	CalAmount float64
} {
	// Sort the r.Currency slice by the Currency field
	sort.Slice(r.Currency, func(i, j int) bool {
		return r.Currency[i].Currency < r.Currency[j].Currency
	})

	// Create a new slice to hold the sorted data
	sortedCurrency := make([]struct {
		Currency  string
		CalAmount float64
	}, len(r.Currency))

	ratio := math.Pow(10, float64(2))

	// Copy the sorted data to the new slice
	for i, data := range r.Currency {
		sortedCurrency[i] = struct {
			Currency  string
			CalAmount float64
		}{
			Currency:  data.Currency,
			CalAmount: math.Round(data.CalAmount*ratio) / ratio,
		}
	}

	return sortedCurrency
}

func (r *CalResult) BalanceToDisplay() float64 {
	ratio := math.Pow(10, float64(2))
	return math.Round(r.CalAllBalance*ratio) / ratio
}

func (g *Gateway) ToDisplay() map[string]interface{} {
	return map[string]interface{}{
		"id":           g.ID,
		"name":         g.Name,
		"display_name": g.DisplayName,
	}
}
