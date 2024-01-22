/*
	@author: Sushant
*/

package http_handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	es "github.com/sushant102004/CatalogIQ/internal/elasticsearch"
)

type HTTPHandler struct {
	// Elasticsearch dependency
	es *es.ES
}

func NewHTTPHandler(es *es.ES) *HTTPHandler {
	return &HTTPHandler{
		es: es,
	}
}

func (h *HTTPHandler) HandleIndexData(ctx *fiber.Ctx) error {
	// This function will first create indexes if doesn't exist and then index the data in all the created indexes

	var body interface{}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(
			map[string]string{
				"message": "Invalid input data",
				"error":   err.Error(),
			},
		)
	}

	doc := body.(map[string]interface{})

	indexes := []string{}

	for k := range doc {
		indexes = append(indexes, k)
	}

	if err := h.es.CreateIndex(indexes); err != nil {
		return ctx.Status(500).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	if err := h.es.IndexItem(indexes, body); err != nil {
		return ctx.Status(500).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(
		map[string]string{
			"message": "Data indexed successfully",
		},
	)
}

func (h *HTTPHandler) HandleSearchDocument(ctx *fiber.Ctx) error {
	// Taking search query in query params for now

	var body interface{}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(
			map[string]string{
				"message": "Invalid input data",
				"error":   err.Error(),
			},
		)
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return ctx.Status(400).JSON(
			map[string]string{
				"message": "Unable to marshal search query",
				"error":   "Internal Server Error" + err.Error(),
			},
		)
	}

	resp, err := h.es.SearchDocument(string(bodyBytes))
	if err != nil {
		return ctx.Status(400).JSON(
			map[string]string{
				"message": "Unable to perform search",
				"error":   "Internal Server Error" + err.Error(),
			},
		)
	}

	ctx.Status(200).JSON(map[string]interface{}{
		"message": "Search Successfull",
		"data":    resp,
	})

	return nil
}
