package main

import (
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

	var products []eureka.Product
	err := client.Products().List(&products)
	if err != nil {
		slog.Error("ListProducts error:", err)
		os.Exit(1)
	}

	for _, product := range products {
		var ideas []eureka.Idea
		var releases []eureka.Release

		err = client.Ideas().For(&product).List(&ideas)

		if err != nil {
			slog.Error("ListIdeas error:", err)
		}

		err = client.Releases().For(&product).List(&releases)
		if err != nil {
			slog.Error("ListReleases error:", err)
		}
	}
}
