package main

import (
	"encoding/json"
	"github.com/zenful-ai/eureka"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	apiToken, ok := os.LookupEnv("AHA_API_KEY")
	if !ok {
		slog.Error("AHA_API_KEY environment variable is not set")
		os.Exit(1)
	}

	request, err := http.NewRequest(http.MethodGet, "https://zenful-ai.aha.io/api/v1/products", nil)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	request.Header.Add("Authorization", "Bearer "+apiToken)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	var productListResponse eureka.ProductListResponse
	b, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(b, &productListResponse)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
