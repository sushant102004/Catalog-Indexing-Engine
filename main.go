/*
	@author: Sushant
*/

package main

import (
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	http_handler "github.com/sushant102004/CatalogIQ/http/handlers"
	es "github.com/sushant102004/CatalogIQ/internal/elasticsearch"
)

var (
	ElasticSearchEndpoint = ""
	APIKey                = ""
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load godotenv: " + err.Error())
	}

	endpoint := os.Getenv("ElasticSearchEndpoint")
	apiKey := os.Getenv("ElasticSearchAPIKey")

	if endpoint == "" || apiKey == "" {
		panic("ElasticSearchEndpoint or APIKey not set")
	}

	ElasticSearchEndpoint = endpoint
	APIKey = apiKey
}

func main() {
	cnf := elasticsearch.Config{
		Addresses: []string{ElasticSearchEndpoint},
		APIKey:    APIKey,
	}

	esClient, err := elasticsearch.NewClient(cnf)
	if err != nil {
		panic(err)
	}

	client := es.NewESClient(esClient)

	app := fiber.New()

	handler := http_handler.NewHTTPHandler(client)

	app.Post("/index-data", handler.HandleIndexData)
	app.Get("/search", handler.HandleSearchDocument)

	app.Listen(":5000")
}
