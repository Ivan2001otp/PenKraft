package controllers

import (
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request){

	if (r.Method != http.MethodGet) {
		http.Error(w, "Invalid request. Supposed to be GET request", http.StatusMethodNotAllowed)
		return;
	}
}