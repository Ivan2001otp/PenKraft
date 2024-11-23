package controllers

import (
	"PencraftB/repository"
	"PencraftB/utils"
	"PencraftB/models"
	"github.com/go-playground/validator/v10"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
	"github.com/gorilla/mux"
)


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
	tag.Is_delete=false;
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

// fetch tag by id (GET)
func FetchTagController(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodGet){
		utils.GetErrorResponse(w, http.StatusMethodNotAllowed, "supposed to be GET request !")
		return;
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 80 * time.Second);
	defer cancel();

	vars := mux.Vars(r);
	tag_id := vars["tag_id"];
	log.Println("Fetching Tag Id : ",tag_id)

	mongoDb := repository.GetMongoDBClient();

	tag,err := mongoDb.FetchTagbyId(ctx, utils.ALL_TAG, tag_id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError);
		log.Printf("Something went wrong while fetching tag by tagid : %v",err)
		return;
	}

	log.Println("Fetched tag by id is successful !")
	utils.GetSuccessResponse(w, http.StatusOK)
	json.NewEncoder(w).Encode(status{
		"message":"success",
		"status":http.StatusOK,
		"data":tag,
	})
	
}

func SoftDeleteTagController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		utils.GetErrorResponse(w, http.StatusMethodNotAllowed, "supposed to be DELETE request")
		return;
	}

	vars := mux.Vars(r)
	tag_id := vars["tag_id"];

	var ctx, cancel = context.WithTimeout(context.Background(), 70 * time.Second)
	defer cancel();

	mongoDb := repository.GetMongoDBClient();

	err  := mongoDb.SoftDeleteTagbyId(ctx, tag_id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	utils.GetSuccessResponse(w, http.StatusOK);
	json.NewEncoder(w).Encode(
		status{
			"message":"soft delete success",
			"status":http.StatusOK,
			"data":tag_id,
		},
	)

}

// method handler for hard-delete
func HardDeleteController(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodDelete){
		utils.GetErrorResponse(w, http.StatusMethodNotAllowed, "supposed to be DELETE request")
		return;
	}

	vars := mux.Vars(r);
	tag_id := vars["tag_id"]
	log.Println("Deleting tag with tagid : ",tag_id)

	var ctx, cancel = context.WithTimeout(context.Background(), 80 * time.Second);
	
	defer cancel();

	mongoDb := repository.GetMongoDBClient();
	err := mongoDb.HardDeleteTagbyId(ctx, tag_id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	utils.GetSuccessResponse(w, http.StatusOK)
	json.NewEncoder(w).Encode(
		status{
			"message":"success",
			"status":http.StatusOK,
			"data":tag_id,
		},
	)
}