package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-portal/internal/domain/exchange_rate"
	"payment-portal/internal/domain/gateway"
	"payment-portal/internal/domain/transaction"
	"payment-portal/internal/middleware"
	"time"
)

func gatewaysRoutes(
	router *gin.Engine,
	mg *middleware.Middleware,
	gatewayRepository *gateway.Repository,
	exchangeRateRepository *exchange_rate.Repository,
	transactionsRepository *transaction.Repository,
) {
	router.GET("/api/portal/gateway/v1/list", mg.AuthToken(), func(c *gin.Context) {
		list := gatewayRepository.GetAllActive()

		var formattedList []map[string]interface{}
		for _, each := range list {
			formattedList = append(formattedList, each.ToDisplay())
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Gateways list",
			"data": gin.H{
				"gateways": formattedList,
			},
		})
	})

	router.GET("/api/portal/gateway/v1/yesterday-eod", mg.AuthToken(), func(c *gin.Context) {
		total := transactionsRepository.YesterdayNumEodTransactions()
		exchangeRates := exchangeRateRepository.GetAll()
		yesterday := time.Now().AddDate(0, 0, -1)
		list := gatewayRepository.GetAllWithEod(yesterday)

		calResult := gateway.CalGateways(list, exchangeRates)

		c.JSON(http.StatusOK, gin.H{
			"message": "Yesterday Eod Dashboard Gateways.",
			"data": gin.H{
				"data": gin.H{
					"total_transactions": total,
					"all_balance": gin.H{
						"cal_all_balance": calResult.BalanceToDisplay(),
						"currency":        calResult.CurrencyToDisplay(),
					},
				},
			},
		})
	})

}
