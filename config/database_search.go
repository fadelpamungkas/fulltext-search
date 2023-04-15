package config

import (
	"log"
	"net/http"

	"github.com/meilisearch/meilisearch-go"
)

type DatabaseSearchConfig struct {
	Host   string
	APIKey string
	Index  string
}

func NewDBSearchIndex(config DatabaseSearchConfig) (*meilisearch.Index, error) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   config.Host,
		APIKey: config.APIKey,
	})

	var index *meilisearch.Index
	index, err := client.GetIndex(config.Index)

	if err != nil {
		if isIndexNotFoundError(err) {
			// Create the index if it does not exist
			task, createErr := client.CreateIndex(&meilisearch.IndexConfig{
				Uid:        config.Index,
				PrimaryKey: "id",
			})
			if createErr != nil {
				return nil, createErr
			}

			log.Println("Meilisearch index created: ", task.IndexUID)

		} else {
			return nil, err
		}
	}

	return index, nil
}

func isIndexNotFoundError(err error) bool {
	if meiliErr, ok := err.(*meilisearch.Error); ok {
		return meiliErr.StatusCode == http.StatusNotFound
	}
	return false
}
