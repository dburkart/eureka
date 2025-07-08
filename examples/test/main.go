package main

import (
	"fmt"
	"github.com/dburkart/eureka"
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
		var ideas []eureka.Idea

		err = client.ListIdeas(&eureka.ListIdeasOptions{
			ProductId: &product.ID,
		}).ForEach(func(response *eureka.IdeaListResponse) {
			ideas = append(ideas, response.Ideas...)
		})
		if err != nil {
			slog.Error("ListIdeas error:", err)
		}
	}

	fmt.Println(response)
}
