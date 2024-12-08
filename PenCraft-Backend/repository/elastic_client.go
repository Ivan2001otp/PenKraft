package repository

import (
	"PencraftB/models"
	"PencraftB/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
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


// save the blog 
func  SaveBlogToES(blog models.Blog) {

	esClient = GetElasticsearchClient();
	
	blogJson, err:= json.Marshal(blog)

	if err != nil {
		log.Println("Error occured while marshalling process.(SaveBlogToES)");
		log.Printf("%v",err);
		return;	
	}

	// indexing(saving) the data
	req := esapi.IndexRequest{
		Index: utils.BLOG_COLLECTION,
		Body: bytes.NewReader(blogJson),
		DocumentID: blog.Blog_id,
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		log.Fatalf("Error getting response (SaveBlogToES). %v",err)
		return;
	}

	defer res.Body.Close()
	log.Println("The response on saving to ES : ",res)
	log.Println("Successfully saved to ES.")
}

func DeleteBlogToES(blog models.Blog) error{
	esClient = GetElasticsearchClient()

	res, err := esClient.Delete	(utils.BLOG_COLLECTION, blog.Blog_id)
	if err != nil {
		return fmt.Errorf("error deleting document: %v",err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from es: %s",res.String())
	}

	fmt.Println("Document with ID %S is deleted.")
	return nil;
}

func SearchBlogByTitleorExcerpt(titleField string, excerptField string) ([]*models.Blog, error) {
	if utils.IsEmpty(excerptField) && utils.IsEmpty(titleField) {
		return nil, fmt.Errorf("Neither title or excerpt is given to search blog !");
	}


}

func searchBlogByTitle(titleField string) []*models.Blog{

}

func searchBlogByexcerpt(excerpt string) []*models.Blog {
	var query map[string]interface{}
	esClient = GetElasticsearchClient();

	if excerpt == utils.EXCERPT{
		query = map[string]interface{}{
			"query":map[string]interface{}{
				"match":map[string]interface{}{
					utils.EXCERPT:  excerpt,
				},
			},
		}
	} else {
		log.Println("Invalid queryField provided : ",excerpt);
		return nil;
	}

	queryJson, err := json.Marshal(query)
	if err != nil {
		log.Println("Could not marshal . (searchBlogByexcerpt) -> ",err)
		return nil;
	}

	res, err := esClient.Search(
		esClient.Search.WithIndex(utils.BLOG_COLLECTION),
		esClient.Search.WithBody(bytes.NewReader(queryJson)),
		esClient.Search.WithPretty(),
	)
	if err != nil {
		log.Println("Something went wrong(searchBlogByexcerpt)->",err);
		return nil;
	}

	defer res.Body.Close();


	if res.IsError() {
		log.Println("something wrong in response shit (searchBlogbyExcerpt)")
		log.Println(res.String())
		return nil;
	}


	// parse the response body
	var result struct {
		Hits struct {
			Hits[] struct {
				Source models.Blog `json:"_source"`
			}`json:"hits"`
		}`json:"hits"`
	}

}