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
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type status map[string]interface{}

type Response struct {
	Data       []models.Blog `json:"data"`
	TotalCount int           `json:"total_count"`
}

// fetches all the available tags for the blogs (GET)
func FetchAllTagController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request type. Supposed to be GET request", http.StatusBadRequest)
		return
	}

	mongoDb := repository.GetMongoDBClient()
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

// creates tags (POST)
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
	mongoDb := repository.GetMongoDBClient()
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

// creates blog (POST)
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
		log.Printf("Error while validating request body %v", validationErr.Error())
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

	err = redisClient.SetinHash(context.Background(), op)
	if err != nil {
		log.Printf("Could not store data in redis in blogController %v", err)
		w.Write([]byte("Failed to save in redis cache"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// push the data to ms queue
	// save data in redis..
	blogId := fmt.Sprintf(blogModel.Blog_id)
	err = redisClient.PushToMessageQueue(context.Background(), utils.MESSAGE_QUEUE_NAME, blogId)

	if err != nil {
		w.Write([]byte("Failed to push task to MQ!"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	successResponse := models.SuccessResponse{
		ID:      blogModel.Blog_id,
		Message: fmt.Sprintf("New blog %s created", blogModel.Title),
	}

	utils.GetSuccessResponse(w, http.StatusCreated)

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

	var ctx, cancel = context.WithTimeout(context.Background(), 80*time.Second)
	var listOfBlog []models.Blog
	mongoDb := repository.GetMongoDBClient()
	redisDb := repository.GetRedisInstance()

	// First check the data in redis
	listOfBlog, err := redisDb.FetchAllBlogfromRedis(ctx)

	if err != nil {
		defer cancel()
		log.Println("Error on fetching data from redis (FetchAllBlogController)")
		utils.GetErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if listOfBlog != nil && len(listOfBlog) > 0 {
		defer cancel()
		//if the data is present in redis, return it.
		log.Println("Data fetched from redis !")
		utils.GetSuccessResponse(w, http.StatusOK)

		json.NewEncoder(w).Encode(status{
			"status": http.StatusOK,
			"data":   listOfBlog,
		})

		return
	}

	// Cache-Miss obtained, read the database to get the demanded data.
	listOfBlog, err = mongoDb.FetchAllBlogs()
	if err != nil {
		defer cancel()
		log.Fatalf("Something went wrong while fetching from mongo(FetchAllBlogController) : %v", err)
		return
	}

	if len(listOfBlog) > 0 {
		defer cancel()

		err = redisDb.DeleteDatafromRedisHashset(ctx, utils.BLOG_COLLECTION, listOfBlog)
		if err != nil {
			log.Println("Failed to delete cache in redis (FetchAllBlogController)")
			log.Fatalf("Error : %v", err)
			return
		}

		// Now the data from MongoDB, is stored in Redis.
		err = redisDb.SaveAllBlogtoRedis(ctx, listOfBlog)
		if err != nil {
			log.Println("Failed to write the new data to redis cache(FetchAllBlogController)!")
			log.Fatalf("Error : %v", err)
			return
		}
	}
	// flush out the old data

	defer cancel()
	utils.GetSuccessResponse(w, http.StatusOK)
	json.NewEncoder(w).Encode(listOfBlog)
}


// fetch blog by blogId (GET)
func FetchBlogbyBlogIdController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		utils.GetErrorResponse(w, http.StatusMethodNotAllowed, "supposed to be GET request !")
		return;
	}

	vars := mux.Vars(r)
	blogId := vars["blog_id"];
	log.Println("Fetching blog with blog-id : ",blogId)

	var ctx,cancel = context.WithTimeout(context.Background(), 80 * time.Second)
	defer cancel();

	redisDb := repository.GetRedisInstance();
	mongoDb := repository.GetMongoDBClient();

	blog,err := redisDb.FetchBlogbyBlogid(ctx, blogId, utils.BLOG_COLLECTION);	

	if err != nil {
		log.Println("Could not fetch blog by blogid in BlogController -> FetchBlogbyBlogIdController()");
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}

	// if the response is in redis, fetch it from main memory itself.(Cache Hit case.)
	if blog != nil {
		
		utils.GetSuccessResponse(w, http.StatusAccepted)
		json.NewEncoder(w).Encode(
			status{
				"message":"success",
				"status":http.StatusAccepted,
				"data":blog,
			},
		)

		return;
	}

	// Cache Miss(if the response is not present in redis,fetch from redis and cache the same.)
	blog,err = mongoDb.FetchBlogbyBlogId(ctx, utils.BLOG_COLLECTION, blogId)
	log.Println(blog);

	if err != nil {
		utils.GetErrorResponse(w, http.StatusInternalServerError, err.Error())
		return;
	}
	// caching in redis
	log.Println("Caching in redis")
	
	var tempList[] models.Blog
	tempList = append(tempList, *blog)
	err = redisDb.SaveAllBlogtoRedis(ctx, tempList)

	if err != nil {
		log.Println("Failed to cache the data to redis")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}

	json.NewEncoder(w).Encode(
		status{
			"message":"success",
			"status":http.StatusAccepted,
			"data":*blog,
		},
	)

}

// we can use this controller to update the body/ softdelete(PUT)
func UpdateBlogController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		utils.GetErrorResponse(w, http.StatusBadRequest, "supposed to be PUT request !")
		return
	}

	var requestPayload models.Blog
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		log.Println(err.Error())
	}

	blog_id := requestPayload.Blog_id
	log.Println("the request body is ", requestPayload)
	log.Println("Blog id to fetched is ", blog_id)

	redisDb := repository.GetRedisInstance()
	mongoDb := repository.GetMongoDBClient()

	var ctx, cancel = context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	blog, err := redisDb.FetchBlogbyBlogid(ctx, blog_id, utils.BLOG_COLLECTION)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	// updating data
	if requestPayload.Body != "" {
		blog.Body = requestPayload.Body
	}

	if requestPayload.Excerpt != "" {
		blog.Excerpt = requestPayload.Excerpt
	}

	if requestPayload.Image != "" {
		blog.Image = requestPayload.Image
	}

	if requestPayload.Slug != "" {
		blog.Slug = requestPayload.Slug
	}

	if requestPayload.Title != "" {
		blog.Title = requestPayload.Title
	}

	blog.Is_deleted = requestPayload.Is_deleted

	blog.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// save the updated data to mongodb
	err = mongoDb.UpdateBlog(utils.BLOG_COLLECTION, *blog)
	if err != nil {
		log.Println("Blog Controller -> UpdateBlogbyBlogid -> mongoDb.UpdateBlog()")
		utils.GetErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// save the updated data to redis.
	operationModel := models.Operation{
		Operation_type: "*",
		Data:           *blog,
	}

	// preparing data to delete
	var oldData []models.Blog
	oldData = append(oldData, *blog)
	err = redisDb.DeleteDatafromRedisHashset(ctx, utils.BLOG_COLLECTION, oldData)
	if err != nil {
		log.Println("Blog Controller -> UpdateBlogbyBlogid -> redisDb.DeleteDatafromRedisHashset()")
		utils.GetErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = redisDb.SetinHash(ctx, operationModel)
	if err != nil {
		log.Println("Blog Controller -> UpdateBlogbyBlogid -> redisDb.Set()")
		utils.GetErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println("saved to redis successfully !")

	err = mongoDb.UpdateBlog(utils.BLOG_COLLECTION, *blog)
	if err != nil {
		log.Println("Blog Controller -> UpdateBlogbyBlogid -> mongoDb.UpdateBlog()")
		utils.GetErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println("saved to mongodb successfully !")

	utils.GetSuccessResponse(w, http.StatusOK)
	json.NewEncoder(w).Encode(status{
		"message": fmt.Sprintf("Blog %s updated", blog_id),
		"status":  http.StatusOK,
		"data":    blog,
	})
}

//------********************----------------------********************

// DANGER function ,that deletes all the data(DELETE)
func DeleteAllDataController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		utils.GetErrorResponse(w, http.StatusBadRequest, "Supposed to be DELETE !")
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 80*time.Second)
	redisDb := repository.GetRedisInstance()

	defer cancel()

	err := redisDb.DeleteFromHashset(ctx, utils.BLOG_COLLECTION)
	if err != nil {
		log.Println("Something went wrong while deleting all data from redis(controllers->blogController->DeleteFromHashset)")
		log.Fatalf("Wrong while deleting data from redis : %v ", err)
		return
	}
	log.Println("Deleted all data from redis !")

	go func() {
		mongoDb := repository.GetMongoDBClient()

		// deleting blog data alone
		err := mongoDb.DeleteAllBlogs(utils.BLOG_COLLECTION)
		if err != nil {
			log.Println("Deleting all data went wrong in mongodb !")
		} else {
			log.Println("Deleting all blog data went successfull !")
		}

		// deleting blog - tag relation
		err = mongoDb.DeleteAllRelations(utils.BLOG_R_TAG)
		if err != nil {
			log.Println("Could not delete all relations.BlogController -> DeleteAllDataController()")
		} else {
			log.Println("Successfully deleted all the relations")
		}

	}()

	utils.GetSuccessResponse(w, http.StatusOK)
	json.NewEncoder(w).Encode(
		status{
			"message": "All Blogs deleted",
			"status":  http.StatusOK,
			"data":    "All Blogs deleted",
		},
	)
}
//-------******************-------------------------******************


// Delete specific Blog (DELETE)
func HardDeleteBlogbyBlogidController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		utils.GetErrorResponse(w, http.StatusBadRequest, "Supposed to be DELETE !")
		return
	}

	vars := mux.Vars(r)
	blog_id := vars["blog_id"]
	log.Println("Blog id to be deleted : ",blog_id)

	mongoDb := repository.GetMongoDBClient()
	redisDb := repository.GetRedisInstance();

	var ctx, cancel = context.WithTimeout(context.Background(), 80 * time.Second)
	defer cancel();


	// removing data from main memory.
	blog,err := redisDb.FetchBlogbyBlogid(ctx, blog_id, utils.BLOG_COLLECTION)
	if err != nil {
		log.Println("HardDeleteBlogByBlogidController-> redisDb.FetchBlogbyBlogid()")
		log.Println("Could not fetch blog from redis by blog-id");
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	var tempList []models.Blog
	tempList = append(tempList, *blog)
	err  = redisDb.DeleteDatafromRedisHashset(ctx, utils.BLOG_COLLECTION,tempList)

	if err != nil{
		log.Println("HardDeleteBlogByBlogidController-> redisDb.DeleteDatafromRedisHashset()")
		log.Println("Could not delete blog from redis by blog-id");
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	// remove data from secondary memory mongoDB.
	err = mongoDb.DeleteBlogbyId(utils.BLOG_COLLECTION, blog_id)
	if err != nil {
		log.Println("HardDeleteBlogByBlogidController->DeleteBlogbyId()")
		log.Println("failed to delete blog by blogid from mongodb")
		utils.GetErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.GetSuccessResponse(w, http.StatusOK)
	json.NewEncoder(w).Encode(
		status{
			"message": fmt.Sprintf("Blog %s is deleted permanently", blog_id),
			"status":  http.StatusOK,
			"data":    blog_id,
		},
	)
}
