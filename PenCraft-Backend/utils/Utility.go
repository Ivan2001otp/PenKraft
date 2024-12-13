package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)


var TTL time.Duration = 60 * time.Minute;

type status map[string]interface{}

func GetCollectionByName(collectionName string) string{
	
	switch (collectionName) {
	case Fps_tag:
		log.Println(Fps_tag)
		return FPS_COLLECTION;

	case Sony_tag:
		log.Println(Sony_tag)
		return SONY_COLLECTION;
	
	case Rpg_tag:
		log.Println(Rpg_tag)
		return RPG_COLLECTION;

	case Ps5_tag:
		log.Println(Ps5_tag)
		return Ps5_tag;
	}

	log.Println(collectionName);
	return Ps5_tag;
}

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