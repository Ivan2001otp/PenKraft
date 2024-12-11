package controllers

import (
	"PencraftB/models"
	"PencraftB/repository"
	"PencraftB/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request. Supposed to be GET request", http.StatusMethodNotAllowed)
		return
	}
}

// for testing
func GetAllBlogES(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request. Supposed to be GET request.", http.StatusMethodNotAllowed)
		return
	}

	var blogList []models.Blog

	blogList = *repository.FetchAllBlogFromES()
	if blogList == nil {
		utils.GetErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Blog list is %v", blogList))
		return
	}

	utils.GetSuccessResponse(w, http.StatusOK)
	json.NewEncoder(w).Encode(
		status{
			"message": "Success",
			"status":  http.StatusOK,
			"data":    blogList,
		},
	)
}

func SoftDeleteBlog(w http.ResponseWriter, r *http.Request) {

}
