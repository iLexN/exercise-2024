package gateway

import (
	"payment-portal/internal/domain/exchange_rate"
	"sort"
	"time"
)

func CalGateways(gateways []Gateway, exchangeRates []exchange_rate.ExchangeRate) *CalResult {

	var allBalance float64
	allCurrency := make(map[string]struct {
		Currency  string
		CalAmount float64
	})
	var newGateways []Summary
	var maxLastUpdatedAt time.Time

	for _, gateway := range gateways {

		eachGateway := Summary{
			Gateway:     gateway,
			LastUpdated: time.Now(),
		}

		for _, balance := range gateway.Balances {

			if _, ok := allCurrency[balance.Currency]; !ok {
				allCurrency[balance.Currency] = struct {
					Currency  string
					CalAmount float64
				}{
					Currency:  balance.Currency,
					CalAmount: balance.GetCalAmount(),
				}
			} else {
				currencyData := allCurrency[balance.Currency]
				currencyData.CalAmount += balance.GetCalAmount()
				allCurrency[balance.Currency] = currencyData
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

func mapToSlice(inputMap map[string]struct {
	Currency  string
	CalAmount float64
}) []struct {
	Currency  string
	CalAmount float64
} {
	result := make([]struct {
		Currency  string
		CalAmount float64
	}, 0, len(inputMap))

	for _, data := range inputMap {
		result = append(result, struct {
			Currency  string
			CalAmount float64
		}{
			Currency:  data.Currency,
			CalAmount: data.CalAmount,
		})
	}

	return result
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
