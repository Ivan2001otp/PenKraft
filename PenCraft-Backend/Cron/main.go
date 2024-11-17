package main

import (
	db "PencraftB/repository"
	"PencraftB/utils"
	"context"
	"log"
	"github.com/robfig/cron/v3"
)

var (
	redisClient *db.RedisClient
	mongoClient *db.DBClient
)

func processQueue(){
	
	for {
		
		blogKey, err := redisClient.PopBlogKeyFromQueue(context.Background(), utils.MESSAGE_QUEUE_NAME);

		if err!=nil {
			continue;
		}

		blogId, err := redisClient.PopBlogdataFromBlogkey(context.Background(), utils.MESSAGE_QUEUE_NAME,*blogKey)
		if err!= nil {
			continue
		}

		log.Printf("%s already processed...",*blogId)
	}
}

func main(){

	log.Println("Cron executing..")

	redisClient = db.GetRedisInstance()
	mongoClient = db.NewDBClient()

	cronScheduler := cron.New()

	cronScheduler.AddFunc("*/2 * * * *",func(){
		log.Println("Executing cron job : processQueue")
		processQueue()
	})

	cronScheduler.Start()

	//keep the main function running...
	select{}
}