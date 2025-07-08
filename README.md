# Eureka

Ergonomic library for the [Aha! API](https://www.aha.io/api).

> [!CAUTION]
> This library is a Work in Progress

## Installing

```shell
go get github.com/dburkart/eureka
```

## Using

Check out the examples directory for further examples, but to give a taste, here's how you would retrieve all products in an account:

```go
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
		CompanyName: "yourcompany",
		ApiKey:      apiToken,
	})

	var products []eureka.Product
	err := client.Products().List(&products)
	if err != nil {
		slog.Error("List error:", err)
		os.Exit(1)
	}
}

```