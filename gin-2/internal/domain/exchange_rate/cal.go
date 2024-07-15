package exchange_rate

func CalByExchangeRate(amount *float64, fromCurrency string, toCurrency string, exchangeRates []ExchangeRate) float64 {

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
