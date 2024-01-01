package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	http_handler "github.com/sushant102004/CatalogIQ/http/handlers"
	es "github.com/sushant102004/CatalogIQ/internal/elasticsearch"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load godotenv: " + err.Error())
	}
}

func main() {
	cnf := elasticsearch.Config{
		Addresses: []string{"https://8fa4a0da88f948278910552683db741e.us-central1.gcp.cloud.es.io"},
		Username:  "elastic",
		Password:  "o2PE0mohdwggfKU4fTKnbTvz",
	}

	esClient, err := elasticsearch.NewClient(cnf)
	if err != nil {
		panic(err)
	}

	client := es.NewESClient(esClient)

	app := fiber.New()

	handler := http_handler.NewHTTPHandler(client)

	app.Post("/create-index", handler.HandleIndexData)

	app.Listen(":5000")
}
