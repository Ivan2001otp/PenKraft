package main

import (
	redisDb "PencraftB/repository"
	mongoDb "PencraftB/repository"
	elasticDb "PencraftB/repository"
	routers  "PencraftB/Routes"
	"net/http"
	"log"
)

func main(){
	log.Println("Started main driver function")

	client := mongoDb.GetMongoDBClient()
	rdb := redisDb.GetRedisInstance()
	_ = elasticDb.GetElasticsearchClient() 

	

	log.Println("Starting server on :8080")
	log.Println("MongoDB is alive !")
	log.Println("redisDB is alive !")
	log.Println("elasticDB is alive !")


	defer client.Close();
	defer rdb.Close();

	router := routers.Router()

	if err:= http.ListenAndServe(":8080", router); err != nil{
		log.Fatal(err);
	}
	
}