package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-portal/internal/domain/gateway"
	"payment-portal/internal/middleware"
)

func gatewaysRoutes(
	router *gin.Engine,
	mg *middleware.Middleware,
	gatewayRepository *gateway.Repository,
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
}
