package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"net/http"
	"payment-portal/internal/domain/uipath"
	"payment-portal/internal/domain/uiprocess"
)

func success(router *gin.Engine, repository *uiprocess.Repository) {
	router.POST("/api/webhooks/ui-path/success", func(c *gin.Context) {
		var payload uipath.WebHookPayload

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		process, err := repository.FindByProcessId(payload.ProcessId())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if !process.Gateway.IsActive() {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Gateway is in-active",
			})
			return
		}

		if !payload.IsSuccessful() {
			// try trigger re-try
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Process Faulted",
			})
			return
		}

		// todo: create batch info

		// get OutputArguments
		outputArguments := payload.Job.OutputArguments

		job, err := uipath.GetGatewayJob(&outputArguments, &process.Gateway)
		if err != nil {
			return
		}

		const redisAddr = "127.0.0.1:6379"
		client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
		defer client.Close()

		err = job.Enqueue(client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong1",
			"p":       job,
		})
	})
}
