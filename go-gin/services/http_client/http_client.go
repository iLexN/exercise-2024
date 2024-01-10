package http_client

import (
	"net/http"
	"time"
)

var HttpClient = httpClient()

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}
