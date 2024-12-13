package main

import (
	"PencraftB/models"
	relations "PencraftB/models/Relations"
	db "PencraftB/repository"
	"PencraftB/utils"
	"context"
	"encoding/json"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

var (
	redisClient *db.RedisClient
	mongoClient *db.DBClient
)

func processQueue() {

	for {

		blogKey, err := redisClient.PopBlogKeyFromQueue(context.Background(), utils.MESSAGE_QUEUE_NAME)
		
		if err != nil {
			continue
		}

		var blogId string = ""

		if len(*blogKey) == 0 {
			log.Println("no tasks in queue")
			return;
		}

		for _, item := range *blogKey {
			blogId = item
			log.Println("Key -> ", blogId)

			blogData, err := redisClient.PopBlogdataFromBlogkey(context.Background(), utils.MESSAGE_QUEUE_NAME, blogId)

			if err != nil {
				continue
			}

			var operation models.Operation
			err = json.Unmarshal([]byte(*blogData), &operation)

			if err != nil {
				log.Println("Failed to unmarshall blog data at cron")
				log.Println(err.Error())
				continue
			}

			switch operation.Operation_type {
			case utils.CREATE_OPS:
				log.Printf("Saving blog %s in db....", operation.Data.Blog_id)

				log.Println("Create operation initialized")
				err := saveBlogByCategoryName(operation.Data, mongoClient);

				if err != nil {
					log.Println("Error on saving blog by category !");
					log.Println(err.Error());
					break;
				}

				var relation relations.R_Tag_Blog
				relation.Blog_id = operation.Data.Blog_id
				relation.Tag_id = operation.Data.Tag_id

				mongoClient.SaveRelation(utils.BLOG_R_TAG, relation)
				mongoClient.UpdateTagcount(true, operation.Data.Tag_id);
				break

			case utils.DELETE_OPS:
				log.Println("Delete operation initialized")
				break

			case utils.UPDATE_OPS:
				log.Println("Update operation initialized")
				break

			default:
				log.Println("Invalid operation type found ", operation.Operation_type)
			}

			log.Printf("Data already processed...")
		}

	}
}

func saveBlogByCategoryName(blog models.Blog, mongoClient *db.DBClient) error {

	category := blog.Tag_name
	var err error;

	switch (category) {
	
	case utils.FPS_COLLECTION:
		_, err = mongoClient.SaveBlog(blog, utils.FPS_COLLECTION)
		break;

	case utils.RPG_COLLECTION: 
		_, err = mongoClient.SaveBlog(blog, utils.RPG_COLLECTION)
		break;

	case utils.SONY_COLLECTION:
		_, err = mongoClient.SaveBlog(blog, utils.SONY_COLLECTION)
		break;
	
	case utils.PS5_COLLECTION:
		_, err = mongoClient.SaveBlog(blog, utils.PS5_COLLECTION)
		break;
	
	}

	return err;
}

func flushAllDataFromHashSet() {
	for {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		err := redisClient.CleanSlateonCache(ctx, utils.BLOG_COLLECTION)
		defer cancel()

		if err != nil {
			log.Fatalf("Something wrong while delete all data from cache(flush.go) : %v", err)
			continue
		}
	}
}

func main() {

	log.Println("Cron executing..")

	redisClient = db.GetRedisInstance()
	mongoClient = db.GetMongoDBClient()
	cronScheduler := cron.New()

	cronScheduler.AddFunc("*/50 * * * *", func() {
		log.Println("Executing cron job for flushing cache.")
		// flushAllDataFromHashSet()
	})

	// cronScheduler.AddFunc("*/1 * * * *", func() {
	// 	log.Println("Executing cron job processQueue")
	// 	processQueue()
	// })
	processQueue()

	cronScheduler.Start()

	//keep the main function running...
	select {}
}
