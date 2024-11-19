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

// fetches all the available tags for the blogs
func FetchAllTagController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request type. Supposed to be GET request", http.StatusBadRequest)
		return
	}

	mongoDb := repository.NewDBClient()
	bsonArray, err := mongoDb.FetchAllTags()

	if err != nil {
		log.Println("Error while fetching from DB !")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(bsonArray)
}

// creates tags
func CreateTagController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request type. Supposed to be POST request", http.StatusBadRequest)
		return
	}

	var tag models.Tag
	validationController := validator.New()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tag)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request body"), http.StatusBadRequest)
		log.Println("Failed to parse Tag request in tagcontroller")
		return
	}

	validationErr := validationController.Struct(tag)

	if validationErr != nil {
		w.Write([]byte("validation error persists when creating tags"))
		http.Error(w, "validation errors", http.StatusBadRequest)
		return
	}

	if tag.Tag_name == "" {
		http.Error(w, "Required field tagname not given !", http.StatusInternalServerError)
		return
	}

	tag.ID = primitive.NewObjectID()
	tag.Tag_id = tag.ID.Hex()
	tag.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	tag.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// save the tag to db directly
	mongoDb := repository.NewDBClient()
	result, err := mongoDb.SaveTagOnly(utils.ALL_TAG, tag)
	log.Printf("The result after saving to DB is %v", result)

	if err != nil {
		log.Println("Something went wrong while saving TAG to db !")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	successResponse := models.SuccessResponse{
		ID:      tag.Tag_id,
		Message: fmt.Sprintf("New Tag -> %s created", tag.Tag_name),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(successResponse); err != nil {
		log.Println("Failed to encode the success response in TagController")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// creates blog
func CreateBlogController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request. Supposed to be POST request!", http.StatusMethodNotAllowed)
		return
	}

	var blogModel models.Blog
	validateController := validator.New()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&blogModel)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request Body"), http.StatusBadRequest)
		log.Println("Failed to decode the body in createBlog controller")
		log.Fatal(err)
	}

	//check validation on fields.
	validationErr := validateController.Struct(&blogModel)
	if validationErr != nil {
		w.Write([]byte("Validation on request fields failed"))
		http.Error(w, "Require fields missing or mistyped !", http.StatusBadRequest)
		log.Fatalf("Error while validating request body %v", validationErr.Error())
	}

	blogModel.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	blogModel.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	blogModel.ID = primitive.NewObjectID()
	blogModel.Blog_id = blogModel.ID.Hex()

	var redisClient repository.RedisClient = *repository.GetRedisInstance()

	op := models.Operation{
		Operation_type: utils.CREATE_OPS,
		Data:           blogModel,
	}

	// save data in redis..
	redisKey := fmt.Sprintf(blogModel.Blog_id)

	// convert the "op" to the slice of bytes . Redis only accepts string or bytes
	bytes, err := json.Marshal(op)
	if err != nil {
		log.Println(err.Error())
		log.Println("Failed to convert blog model to slice of bytes in Blog controller.")
		return
	}

	err = redisClient.Set(context.Background(), redisKey, bytes, utils.TTL)

	if err != nil {
		log.Printf("Could not store data in redis in blogController %v", err)
		w.Write([]byte("Failed to save in redis cache"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// push the data to ms queue
	err = redisClient.PushToMessageQueue(context.Background(), utils.MESSAGE_QUEUE_NAME, redisKey)

	if err != nil {
		w.Write([]byte("Failed to push task to MQ!"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	successResponse := models.SuccessResponse{
		ID:      blogModel.Blog_id,
		Message: fmt.Sprintf("New blog %s created", blogModel.Title),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(successResponse); err != nil {
		log.Println("Could not encode success Response for createBlog controller")
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

// fetch all blogs(GET)
func FetchAllBlogController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Write([]byte("Invalid request. Supposed to be GET request ."))
		http.Error(w, "Make GET request !", http.StatusBadRequest)
		return
	}

	mongoDb := repository.NewDBClient()
	bsonArray, err := mongoDb.FetchAllBlogs()

	if err != nil {
		utils.GetErrorResponse(w, http.StatusInternalServerError, "Could not save records in DB")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(bsonArray)
}
