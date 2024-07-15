package Balance

import (
	"payment-portal/internal/domain/exchange_rate"
	"time"
)

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

	return b.CalAmount != nil
}

func (b *Balance) GetCalAmount() float64 {
	if b.CalAmount == nil {
		return 0
	}

	return *b.CalAmount
}

func (b *Balance) ToUsdCalAmount(exchangeRates []exchange_rate.ExchangeRate) float64 {
	if b.CalAmount == nil {
		return 0
	}
	a := exchange_rate.CalByExchangeRate(b.CalAmount, b.Currency, "USD", exchangeRates)

	return a
}
