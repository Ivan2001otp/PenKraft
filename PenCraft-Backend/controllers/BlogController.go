package controllers

import (
	"PencraftB/models"
	"PencraftB/repository"
	"PencraftB/utils"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type status map[string]interface{}

type operation struct {
	operationType string 	`json:"operation"`
	data		  models.Blog `json:"data"`
}


func createBlog(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w,"Invalid request. Supposed to be POST request!", http.StatusMethodNotAllowed)
		return;
	}

	var blogModel models.Blog;
	validateController := validator.New()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&blogModel)

	if err != nil {
		http.Error(w, fmt.Sprintf(""),http.StatusBadRequest)
		log.Println("Failed to decode the body in createBlog controller")
		log.Fatal(err);
	}

	//check validation on fields.
	validationErr := validateController.Struct(blogModel)
	if validationErr != nil {
		w.Write([]byte("Validation on request fields failed"))
		http.Error(w, "Require fields missing or mistyped !",http.StatusBadRequest)
		log.Fatalf("Error while validating request body %v", err)
	}

	blogModel.Created_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	blogModel.Updated_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	blogModel.ID = primitive.NewObjectID()
	blogModel.Blog_id = blogModel.ID.Hex()

	var redisClient repository.RedisClient = *repository.GetRedisInstance();


	op := operation{
		operationType: "create",
		data: blogModel,
	}

	// save data in redis..
	redisKey := fmt.Sprintf("blog:%s", blogModel.Blog_id)
	err = redisClient.Set(context.Background(),redisKey, op, utils.TTL)
	
	if err!= nil {
		log.Println("Could not store data in redis in blogController")
		w.Write([]byte("Failed to save in redis cache"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}

	// push the data to ms queue
	err = redisClient.PushToMessageQueue(context.Background(), utils.MESSAGE_QUEUE_NAME, redisKey)

	if err != nil {
		w.Write([]byte("Failed to push task to MQ!"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}

	successResponse := models.SuccessResponse{
		ID: blogModel.Blog_id,
		Message: fmt.Sprintf("New blog %s created", blogModel.Title),
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	
	if err := json.NewEncoder(w).Encode(successResponse); err != nil {
		log.Println("Could not encode success Response for createBlog controller")
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err),http.StatusInternalServerError)
	}

}