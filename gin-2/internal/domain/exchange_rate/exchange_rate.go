package exchange_rate

import "time"

type ExchangeRate struct {
	ID           uint64 `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	FromCurrency string
	ToCurrency   string
	Rate         float64 `gorm:"type:decimal(10,6)"`
}
