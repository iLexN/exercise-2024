package gateway

import (
	"gorm.io/datatypes"
	"payment-portal/internal/domain/Balance"
	"payment-portal/internal/utility"
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

func (s *Summary) ToDisplay() map[string]interface{} {
	g := s.Gateway.ToDisplay()
	g["last_updated"] = s.LastUpdated.UTC()

	return g
}

type CalResult struct {
	CalAllBalance float64
	Currency      []*CurrencyAmount
	Gateways      []Summary
}

type CurrencyAmount struct {
	Currency string
	Amount   float64
}

func (r *CalResult) CurrencyToDisplay() []*CurrencyAmount {
	// Sort the r.Currency slice by the Currency field
	sort.Slice(r.Currency, func(i, j int) bool {
		return r.Currency[i].Currency < r.Currency[j].Currency
	})

	// Create a new slice to hold the sorted data
	sortedCurrency := make([]*CurrencyAmount, len(r.Currency))

	// Copy the sorted data to the new slice
	for i, data := range r.Currency {
		sortedCurrency[i] = &CurrencyAmount{
			Currency: data.Currency,
			Amount:   utility.RoundFloat(data.Amount, 2),
		}
	}

	return sortedCurrency
}

func (r *CalResult) BalanceToDisplay() float64 {
	return utility.RoundFloat(r.CalAllBalance, 2)
}

func (r *CalResult) GatewaysToDisplay() []interface{} {
	out := make([]interface{}, 0, len(r.Gateways))
	for _, data := range r.Gateways {
		info := data.ToDisplay()
		out = append(out, info)
	}

	return out
}

func (g *Gateway) ToDisplay() map[string]interface{} {
	return map[string]interface{}{
		"id":            g.ID,
		"name":          g.Name,
		"display_name":  g.DisplayName,
		"client_config": g.ClientConfig,
	}
}
