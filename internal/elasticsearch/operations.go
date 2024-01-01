package es

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8"
	catalog "github.com/sushant102004/CatalogIQ/model"
)

type ES struct {
	client *elasticsearch.Client
}

func NewESClient(client *elasticsearch.Client) *ES {
	return &ES{
		client: client,
	}
}

func (es *ES) CreateIndex(indexes []string) error {
	for _, index := range indexes {
		resp, err := es.client.Indices.Create(index)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
	}
	return nil
}

func (es *ES) IndexGroceryItem(data catalog.GroceryItem) error {
	_, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// for _, index := range indexes {
	// 	_, err := es.client.Index(index, strings.NewReader(string(doc)))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Printf("Indexed: %s with data %s\n", index, doc)
	// }
	return nil
}

func (es *ES) IndexFashionItem(data catalog.GroceryItem) error { return nil }
