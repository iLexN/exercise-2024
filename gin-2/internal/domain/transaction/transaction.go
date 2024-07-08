package transaction

import "time"

type Transaction struct {
	ID            uint64     `gorm:"primaryKey;column:id"`
	GatewaysID    int        `gorm:"column:gateways_id"`
	BatchID       int        `gorm:"column:batch_id"`
	PaymentID     *string    `gorm:"column:payment_id"`
	PaymentType   *string    `gorm:"column:payment_type"`
	Amount        *float64   `gorm:"column:amount"`
	Currency      *string    `gorm:"column:currency"`
	TransactionAt *time.Time `gorm:"column:transaction_at"`
	Status        *string    `gorm:"column:status"`
	Notes         *string    `gorm:"column:notes"`
	CreatedAt     time.Time  `gorm:"column:created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at"`
	UserChangeAt  *time.Time `gorm:"column:user_change_at"`
	ReportEodAt   *time.Time `gorm:"column:report_eod_at"`
	OrderNo       *string    `gorm:"column:order_no"`
}
