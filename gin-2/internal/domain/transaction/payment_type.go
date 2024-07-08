package transaction

type PaymentType string

const (
	Deposit     PaymentType = "deposit"
	Withdraw    PaymentType = "withdraw"
	DepositFee  PaymentType = "deposit_fee"
	WithdrawFee PaymentType = "withdraw_fee"

	Settlement    PaymentType = "settlement"
	SettlementFee PaymentType = "settlement_fee"

	TopUp    PaymentType = "top_up"
	TopUpFee PaymentType = "top_up_fee"
)

func TodayPaymentType() []PaymentType {
	return []PaymentType{Deposit, Withdraw}
}

func EodPaymentType() []PaymentType {
	return []PaymentType{
		Deposit,
		Withdraw,
		Settlement,
		TopUp,
	}
}
