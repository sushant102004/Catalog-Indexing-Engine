package http_handler

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	es "github.com/sushant102004/CatalogIQ/internal/elasticsearch"
	catalog "github.com/sushant102004/CatalogIQ/model"
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

	var body catalog.GroceryItem

	if err := ctx.BodyParser(&body); err != nil {
		ctx.Status(400).JSON(
			map[string]string{
				"error": err.Error(),
			},
		)
		return err
	}

	indexes := []string{}

	fields := reflect.TypeOf(body)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]
		indexes = append(indexes, jsonTag)
	}

	if err := h.es.CreateIndex(indexes); err != nil {
		ctx.Status(500).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	ctx.Status(200).JSON(
		map[string]string{
			"message": "Indexes created successfully",
		},
	)
	return nil
}
