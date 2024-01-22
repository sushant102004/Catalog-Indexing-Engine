/*
	@author: Sushant
*/

package es

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
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

		if resp.StatusCode == http.StatusInternalServerError {
			return errors.New("internal server error from elasticsearch system")
		}

		/*
			Uncomment this for debugging purposes.

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return errors.New("internal server error: unable to read response body")
			}

			fmt.Println("Status Code:", resp.StatusCode)
			fmt.Println("Body:", string(body))
		*/
	}
	return nil
}

func (es *ES) IndexItem(indexes []string, doc interface{}) error {
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"

	for _, index := range indexes {
		jsonDoc, err := json.Marshal(doc)
		if err != nil {
			return errors.New("internal server error: unable to marshal json document: " + err.Error())
		}
		resp, err := es.client.Index(index, strings.NewReader(string(jsonDoc)), es.client.Index.WithHeader(headers))
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusInternalServerError {
			return errors.New("internal server error from elasticsearch system")
		}

		/*
			Uncomment for debugging purposes.

			body, err := io.ReadAll(resp.Body)
					if err != nil {
						return err
					}
					fmt.Println("Status Code:", resp.StatusCode)
					fmt.Println("Body: ", string(body))
		*/
	}
	return nil
}
func (es *ES) getAllIndexes() ([]string, error) {
	resp, err := es.client.Indices.Get([]string{"_all"}, es.client.Indices.Get.WithContext(context.Background()))
	if err != nil {
		return nil, errors.New("error unable to get indices: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return nil, errors.New("error elasticsearch response: " + resp.String())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error unable to get indices: " + err.Error())
	}

	var result map[string]interface{}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	var indices []string

	for index := range result {
		if index[0] != '.' {
			indices = append(indices, index)
		}
	}

	return indices, nil
}

func hashDocument(doc map[string]interface{}) (uint32, error) {
	hash := fnv.New32a()
	jsonBytes, err := json.Marshal(doc)
	if err != nil {
		return 0, errors.New("internal server error: unable to hash" + err.Error())
	}
	hash.Write(jsonBytes)
	return hash.Sum32(), nil
}

func (es *ES) SearchDocument(query string) ([]map[string]interface{}, error) {
	indexes, err := es.getAllIndexes()
	if err != nil {
		return nil, errors.New("error unable to get indices: " + err.Error())
	}

	resp, err := es.client.Search(
		es.client.Search.WithIndex(indexes...),
		es.client.Search.WithBody(strings.NewReader(query)),
	)

	if err != nil {
		fmt.Println("error: ", err.Error())
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("error decoding response: ", err.Error())
		return nil, err
	}

	var response []map[string]interface{}

	uniqueDocumentHashes := make(map[uint32]bool)
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})
		docHash, err := hashDocument(doc)
		if err != nil {
			return nil, err
		}

		if _, exists := uniqueDocumentHashes[docHash]; !exists {
			fmt.Println("Unique Hit:", hit)
			uniqueDocumentHashes[docHash] = true
			response = append(response, doc)
		}
	}

	return response, nil
}
