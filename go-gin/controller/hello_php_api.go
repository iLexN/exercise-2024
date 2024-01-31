package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go1/services/http_client"
	"io"

	//	"go1/services/logger"
	"encoding/json"
	"net/http"
)

func PhpApiCall(c *gin.Context) {
	endpoint := "http://php-hyperf:9501/"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		c.JSON(400, gin.H{"error": "new request error"})
		return
	}

	response, err := http_client.HttpClient.Do(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error sending request to API endpoint"})
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Couldn't parse response body"})
		return
	}

	apiResponse, err := createResponseFromBody(body)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": apiResponse,
	})

}

type apiResponse struct {
	Message string `json:"message"`
}

var (
	errUhOh = errors.New("failed to unmarshal JSON to createResponseFromBody")
)

func createResponseFromBody(body []byte) (*apiResponse, error) {
	var response apiResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, errUhOh
	}

	return &response, nil
}
