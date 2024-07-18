package gateway

import (
	"payment-portal/internal/domain/exchange_rate"
	"time"
)

func CalGateways(gateways []Gateway, exchangeRates []exchange_rate.ExchangeRate) *CalResult {

	var allBalance float64
	allCurrency := make(map[string]*CurrencyAmount)
	var newGateways []Summary
	var maxLastUpdatedAt time.Time

	for _, gateway := range gateways {

		eachGateway := Summary{
			Gateway:     gateway,
			LastUpdated: time.Now(),
		}

		for _, balance := range gateway.Balances {

			adjCurrency := exchange_rate.AdjustCurrency(balance.Currency)
			if _, ok := allCurrency[adjCurrency]; !ok {
				allCurrency[adjCurrency] = &CurrencyAmount{
					Currency: adjCurrency,
					Amount:   balance.GetCalAmount(),
				}
			} else {
				currencyData := allCurrency[adjCurrency]
				currencyData.Amount += balance.GetCalAmount()
				allCurrency[adjCurrency] = currencyData
			}

			if balance.HaveCalAmount() {
				allBalance += balance.ToUsdCalAmount(exchangeRates)
			}
			//eachGateway.Balance = sortCurrency(allCurrency)

			if balance.UpdatedAt.After(maxLastUpdatedAt) {
				maxLastUpdatedAt = balance.UpdatedAt
			}
		}
		eachGateway.LastUpdated = maxLastUpdatedAt
		newGateways = append(newGateways, eachGateway)

	}

	return &CalResult{
		CalAllBalance: allBalance,
		Currency:      mapToSlice(allCurrency),
		Gateways:      newGateways,
	}
}

func mapToSlice(inputMap map[string]*CurrencyAmount) []*CurrencyAmount {
	result := make([]*CurrencyAmount, 0, len(inputMap))

	for _, data := range inputMap {
		result = append(result, &CurrencyAmount{
			Currency: data.Currency,
			Amount:   data.Amount,
		})
	}

	return result
}
