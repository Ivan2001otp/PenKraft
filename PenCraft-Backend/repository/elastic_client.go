package repository

import (
	"log"
	"sync"
	"github.com/elastic/go-elasticsearch/v7"
	"PencraftB/utils"
)


// singleton for elasticsearch client
var (
	esClient *elasticsearch.Client
	once1 sync.Once
)

// initializes the elastic search instance and returns the elastic search client.
func GetElasticsearchClient() *elasticsearch.Client {

	once1.Do(func ()  {
		cfg := elasticsearch.Config{
			Addresses: []string{"http://"+utils.KAFKA_BROKER},
		}

		client, err := elasticsearch.NewClient(cfg)
		if err != nil {
			log.Fatalf("Error creating elasticsearch client: %v", err)
			return;
		}
		esClient = client;
		log.Println("Elasticsearch client initialized")
	})

	return esClient;
}