package main

import (
	redisDb "PencraftB/repository"
	"PencraftB/utils"
	mongoDb "PencraftB/repository"
	routers  "PencraftB/Routes"
	"net/http"
	"log"
)

func main(){
	utils.Logger("Started main driver function")

	client := mongoDb.NewDBClient()
	rdb := redisDb.GetRedisInstance()

	log.Println("Starting server on :8080")

	defer client.Close();
	defer rdb.Close();

	router := routers.Router()

	if err:= http.ListenAndServe(":8080", router); err != nil{
		log.Fatal(err);
	}
	
}