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

	var idea eureka.Idea
	idea.ID = "7522564463417005564"

	c := eureka.Comment{
		Body: "<p>This is a test comment!</p>",
	}

	err := client.Comments().For(&idea).Create(&c)
	if err != nil {
		slog.Error("Error creating comment", err)
		os.Exit(1)
	}

	fmt.Println("Comment created")
}
