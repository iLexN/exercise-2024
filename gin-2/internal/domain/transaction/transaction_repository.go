package transaction

import (
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) TodayNumTransactions() int64 {
	now := time.Now()
	startOfDay := startOfCurrentDay(now)
	endOfDay := endOfCurrentDay(now)

	var count int64

	r.Db.Model(&Transaction{}).
		Where("transaction_at BETWEEN ? AND ?", startOfDay, endOfDay).
		Where("status IN ?", TodayPaymentStatus()).
		Where("payment_type IN ?", TodayPaymentType()).
		Count(&count)

	return count
}

func (r *Repository) YesterdayNumEodTransactions() int64 {
	yesterday := time.Now().AddDate(0, 0, -1)

	var count int64

	r.Db.Model(&Transaction{}).
		Where("report_eod_at BETWEEN ? AND ?", startOfCurrentDay(yesterday), endOfCurrentDay(yesterday)).
		Where("payment_type IN ?", EodPaymentType()).
		Count(&count)

	return count

}

func startOfCurrentDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func endOfCurrentDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 59, t.Location())
}
