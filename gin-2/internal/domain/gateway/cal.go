package gateway

import "payment-portal/internal/domain/exchange_rate"

func calGateways(gateways []Gateway, exchangeRate []exchange_rate.ExchangeRate) {
	var eachCurrency []map[string]interface{}
	var allBalance float64
	allCurrency := make(map[string]map[string]interface{})

	for _, gateway := range gateways {

		for _, balance := range gateway.Balances {
			eachCurrency = append(eachCurrency, map[string]interface{}{
				"currency":       balance.Currency,
				"cal_amount":     balance.CalAmount,
				"gateway_amount": balance.GatewayAmount,
			})

			if _, ok := allCurrency[balance.Currency]; !ok {
				allCurrency[balance.Currency] = map[string]interface{}{
					"currency":   balance.Currency,
					"cal_amount": balance.CalAmount,
				}
			} else {
				allCurrency[balance.Currency]["cal_amount"] = allCurrency[balance.Currency]["cal_amount"].(float64) + *balance.CalAmount
			}

			if balance.HaveCalAmount(exchangeRates) {
				allBalance += balance.toUsdCalAmount(exchangeRates)
			}

		}
	}
}
