package main

import (
	redisDb "PencraftB/repository"
	mongoDb "PencraftB/repository"
	routers  "PencraftB/Routes"
	"net/http"
	"log"
)

func main(){
	log.Println("Started main driver function")

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