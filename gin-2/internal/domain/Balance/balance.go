package Balance

import "time"

type Balance struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time

	GatewaysID       int
	BalanceAt        time.Time
	CalAmount        *float64
	GatewayAmount    *float64
	AvailableBalance *float64
	Currency         string
}

func (b *Balance) HaveCalAmount() bool {
	return true
}
