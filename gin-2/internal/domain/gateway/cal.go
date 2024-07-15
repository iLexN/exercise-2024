package gateway

import (
	"payment-portal/internal/domain/exchange_rate"
	"sort"
	"time"
)

func CalGateways(gateways []Gateway, exchangeRates []exchange_rate.ExchangeRate) *CalResult {
	var eachCurrency []map[string]interface{}
	var allBalance float64
	allCurrency := make(map[string]map[string]interface{})
	var newGateways []Summary
	var maxLastUpdatedAt time.Time

	for _, gateway := range gateways {

		eachGateway := Summary{
			Gateway:     gateway,
			LastUpdated: time.Now(),
		}

		for _, balance := range gateway.Balances {
			eachCurrency = append(eachCurrency, map[string]interface{}{
				"currency":       balance.Currency,
				"cal_amount":     balance.CalAmount,
				"gateway_amount": balance.GatewayAmount,
			})

			if _, ok := allCurrency[balance.Currency]; !ok {
				allCurrency[balance.Currency] = map[string]interface{}{
					"currency":   balance.Currency,
					"cal_amount": balance.GetCalAmount(),
				}
			} else {
				allCurrency[balance.Currency]["cal_amount"] = allCurrency[balance.Currency]["cal_amount"].(float64) + balance.GetCalAmount()
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
		Currency:      mapToSlice(sortCurrency(allCurrency)),
		Gateways:      newGateways,
	}
}

func mapToSlice(inputMap map[string]map[string]interface{}) []map[string]interface{} {
	var outputSlice []map[string]interface{}
	for _, data := range inputMap {
		tempMap := make(map[string]interface{})
		for key, value := range data {
			tempMap[key] = value
		}

		outputSlice = append(outputSlice, tempMap)
	}
	return outputSlice
}

func sortCurrency(allCurrency map[string]map[string]interface{}) map[string]map[string]interface{} {

	// Convert the map keys to a slice for sorting
	var currencies []string
	for currency := range allCurrency {
		currencies = append(currencies, currency)
	}

	// Sort the currencies slice
	sort.Slice(currencies, func(i, j int) bool {
		return allCurrency[currencies[i]]["currency"].(string) < allCurrency[currencies[j]]["currency"].(string)
	})

	// Create a new map to store the sorted currencies
	sortedCurrency := make(map[string]map[string]interface{})

	// Populate the new map with the sorted currencies
	for _, currency := range currencies {
		sortedCurrency[currency] = allCurrency[currency]
	}

	return sortedCurrency

}
