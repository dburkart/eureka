package main

import (
	"fmt"
	"github.com/zenful-ai/eureka"
	"log/slog"
	"os"
)

func main() {
	apiToken, ok := os.LookupEnv("AHA_API_KEY")
	if !ok {
		slog.Error("AHA_API_KEY environment variable is not set")
		os.Exit(1)
	}

	client := eureka.NewAhaClient(&eureka.AhaClientOptions{
		CompanyName: "zenful-ai",
		ApiKey:      apiToken,
	})

	response, err := client.ListProducts().Do()
	if err != nil {
		slog.Error("ListProducts error:", err)
		os.Exit(1)
	}

	for _, product := range response.Products {
		ideas, err := client.ListIdeas(&eureka.ListIdeasOptions{
			ProductId: &product.ID,
		}).Do()
		if err != nil {
			slog.Error("ListIdeas error:", err)
		}

		fmt.Println(ideas)
	}

	fmt.Println(response)
}
