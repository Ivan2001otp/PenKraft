package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)


var TTL time.Duration = 60 * time.Minute;

type status map[string]interface{}

func GetSuccessResponse(w http.ResponseWriter, statusCode int){
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(statusCode)
}

func GetErrorResponse(w http.ResponseWriter,statusCode int, message string) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)

	response := status{
		"statusCode":statusCode,
		"message":message,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil{
		log.Println("unable to create error response in utility.")
		http.Error(w, "Unable to parse response",http.StatusInternalServerError)
	}

}

func IsEmpty(str string) bool {
	return len(str)==0;
}

func NotEmpty(str string) bool {
	return len(str) > 0;
}