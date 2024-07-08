package transaction

type PaymentStatus string

const (
	Completed PaymentStatus = "completed"
	Pending   PaymentStatus = "pending"
	Failed    PaymentStatus = "failed"
)

func TodayPaymentStatus() []PaymentStatus {
	return []PaymentStatus{Completed, Pending}
}
