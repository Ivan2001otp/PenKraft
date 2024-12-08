package repository

import (
	"PencraftB/models"
	"PencraftB/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// singleton for elasticsearch client
var (
	esClient *elasticsearch.Client
	once1    sync.Once
)

// initializes the elastic search instance and returns the elastic search client.
func GetElasticsearchClient() *elasticsearch.Client {

	once1.Do(func() {
		cfg := elasticsearch.Config{
			Addresses: []string{"http://" + utils.KAFKA_BROKER},
		}

		client, err := elasticsearch.NewClient(cfg)
		if err != nil {
			log.Fatalf("Error creating elasticsearch client: %v", err)
			return
		}
		esClient = client
		log.Println("Elasticsearch client initialized")
	})

	return esClient
}

// save the blog
func SaveBlogToES(blog models.Blog) {

	esClient = GetElasticsearchClient()

	blogJson, err := json.Marshal(blog)

	if err != nil {
		log.Println("Error occured while marshalling process.(SaveBlogToES)")
		log.Printf("%v", err)
		return
	}

	// indexing(saving) the data
	req := esapi.IndexRequest{
		Index:      utils.BLOG_COLLECTION,
		Body:       bytes.NewReader(blogJson),
		DocumentID: blog.Blog_id,
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		log.Fatalf("Error getting response (SaveBlogToES). %v", err)
		return
	}

	defer res.Body.Close()
	log.Println("The response on saving to ES : ", res)
	log.Println("Successfully saved to ES.")
}

func DeleteBlogToES(blog models.Blog) error {
	esClient = GetElasticsearchClient()

	res, err := esClient.Delete(utils.BLOG_COLLECTION, blog.Blog_id)
	if err != nil {
		return fmt.Errorf("error deleting document: %v", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from es: %s", res.String())
	}

	fmt.Println("Document with ID %S is deleted.")
	return nil
}

func SearchBlogByTitleorExcerpt(titleField string, excerptField string) (*([]models.Blog), error) {
	if utils.IsEmpty(excerptField) && utils.IsEmpty(titleField) {
		return nil, fmt.Errorf("Neither title or excerpt is given to search blog !")
	}

	var blogList []models.Blog

	if !utils.IsEmpty(titleField) {
		blogList = *searchBlogByTitle(titleField)
	}

	if len(blogList) > 0 {
		return &blogList, nil
	}

	if !utils.IsEmpty(excerptField) {
		blogList = *searchBlogByexcerpt(excerptField)
	}

	return &blogList, nil

}

func searchBlogByTitle(titleField string) *([]models.Blog) {
	var query map[string]interface{}
	esClient = GetElasticsearchClient()

	if titleField == utils.TITLE {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"prefix": map[string]interface{}{
					utils.TITLE: map[string]interface{}{
						"value": titleField,
					},
				},
			},
		}
	} else {
		log.Println("Invalid queryField provided: ", titleField)
	}

	queryJson, err := json.Marshal(query)
	if err != nil {
		log.Println("Could not marshal . (searchBlogByTitle) -> ", err)
		return nil
	}

	res, err := esClient.Search(
		esClient.Search.WithIndex(utils.BLOG_COLLECTION),
		esClient.Search.WithBody(bytes.NewReader(queryJson)),
		esClient.Search.WithPretty(),
	)

	if err != nil {
		log.Println("Something went wrong(searchBlogByTitle)->", err)
		return nil
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Println("something wrong in response shit (searchBlogbyExcerpt)")
		log.Println(res.String())
		return nil
	}

	var result struct {
		Hits struct {
			Hits []struct {
				Source models.Blog `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Println("(searchBlogByexcerpt)Error parsing response: %v", err)
		return nil
	}

	var blogList []models.Blog

	for _, hit := range result.Hits.Hits {
		blogList = append(blogList, hit.Source)
	}

	return &blogList
}

func searchBlogByexcerpt(excerpt string) *([]models.Blog) {
	var query map[string]interface{}
	esClient = GetElasticsearchClient()

	if excerpt == utils.EXCERPT {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"prefix": map[string]interface{}{
					utils.EXCERPT: map[string]interface{}{
						"value": excerpt,
					},
				},
			},
		}
	} else {
		log.Println("Invalid queryField provided : ", excerpt)
		return nil
	}

	queryJson, err := json.Marshal(query)
	if err != nil {
		log.Println("Could not marshal . (searchBlogByexcerpt) -> ", err)
		return nil
	}

	res, err := esClient.Search(
		esClient.Search.WithIndex(utils.BLOG_COLLECTION),
		esClient.Search.WithBody(bytes.NewReader(queryJson)),
		esClient.Search.WithPretty(),
	)
	if err != nil {
		log.Println("Something went wrong(searchBlogByexcerpt)->", err)
		return nil
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Println("something wrong in response shit (searchBlogbyExcerpt)")
		log.Println(res.String())
		return nil
	}

	// parse the response body
	var result struct {
		Hits struct {
			Hits []struct {
				Source models.Blog `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	// return matched blogs whose excerpts are matched.
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Println("(searchBlogByexcerpt)Error parsing response: %v", err)
		return nil
	}

	var blogList []models.Blog

	for _, hit := range result.Hits.Hits {
		blogList = append(blogList, hit.Source)
	}

	return &blogList
}

func updateBlogToDelete(deletedBlog models.Blog) error {
	esClient = GetElasticsearchClient()

	updateDoc := map[string]interface{}{
		"doc":map[string]interface{}{
			"blog_id": deletedBlog.Blog_id,
		},
	}

	updatedDocJson, err := json.Marshal(updateDoc)
	if err != nil {
		log.Println("(updateBlogToDelete)Error in marshalling document : %v", err)
		return err;
	}


	res, err := esClient.Update(
		utils.BLOG_COLLECTION,
		deletedBlog.Blog_id,
		bytes.NewReader(updatedDocJson),
		esClient.Update.WithPretty(),
	)
	if err != nil {
		log.Println("(updateBlogToDelete)Something went wrong while updating data in elasticSearch.")
		return err;
	}


	defer res.Body.Close();
	if res.IsError() {
		return fmt.Errorf("Error response from elasticsearch is %v", res.String())
	}

	log.Printf("Blog with %s successfully soft deleted\n", deletedBlog.Blog_id);
	return nil;
}

func fetchAllBlogFromES() ( interface{} ) {
	var blogList []models.Blog


	query := map[string]interface{}{
		"query":map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}


	queryJson, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("Error marshalling query : %s", err)
	}


	res, err := esClient.Search(
		esClient.Search.WithIndex(utils.BLOG_COLLECTION),
		esClient.Search.WithBody(bytes.NewReader(queryJson)),
		esClient.Search.WithPretty(),
	)
	if err != nil {
		return fmt.Errorf("Error performing search query : %v", err);
	}

	defer res.Body.Close();

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	var result struct {
		Hits struct {
			Hits []struct {
				Source models.Blog `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return fmt.Errorf("error parsing response: %s", err)
	}



	for _, hit := range result.Hits.Hits {
		blogList = append(blogList, hit.Source)
	}

	return &blogList
}
