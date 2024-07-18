package exchange_rate

func CalByExchangeRate(amount *float64, fromCurrency string, toCurrency string, exchangeRates []ExchangeRate) float64 {

	fromCurrency = AdjustCurrency(fromCurrency)

	if amount == nil {
		return 0
	}

	for _, rate := range exchangeRates {
		if rate.FromCurrency == fromCurrency && rate.ToCurrency == toCurrency {
			return *amount * rate.Rate
		}
	}
	return 0
}

func AdjustCurrency(currency string) string {
	if currency == "MYR-FPX" {
		return "MYR"
	}

	return currency
}
