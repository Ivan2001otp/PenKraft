package main

import (
	routers "PencraftB/Routes"
	mongoDb "PencraftB/repository"
	redisDb "PencraftB/repository"
	"log"
	"net/http"
	"github.com/rs/cors"
)

func main(){
	log.Println("Started main driver function")

	client := mongoDb.GetMongoDBClient()
	rdb := redisDb.GetRedisInstance()
	//_ = elasticDb.GetElasticsearchClient() 
	//elasticDb.PingElasticsearch()
	
	log.Println("Starting server on :8080")
	log.Println("MongoDB is alive !")
	log.Println("redisDB is alive !")
	//log.Println("elasticDB is alive !")

	defer client.Close();
	defer rdb.Close();

	//cors options
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	router := routers.Router()
	
	handler := corsOptions.Handler(router)

	log.Println("Server listening to 8080 port !");
	if err:= http.ListenAndServe(":8080", handler); err != nil{
		log.Fatal(err);
	}
	
}