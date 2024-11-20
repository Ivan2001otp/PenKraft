package main

import (
	"PencraftB/utils"
	"context"
	"log"
	"time"
	"github.com/robfig/cron/v3"
)



func flushAllDataFromHashSet() {
	for {
		var ctx, cancel = context.WithTimeout(context.Background(), 80 * time.Second)
		err := redisClient.CleanSlateonCache(ctx,utils.BLOG_COLLECTION);
		defer cancel();
		

		if err != nil {
			log.Fatalf("Something wrong while delete all data from cache(flush.go) : %v",err);
			continue
		}
	}
}

func MainDriver(){
	log.Println("Flush-Cron executing...")

	cronScheduler := cron.New()

	cronScheduler.AddFunc("*/15 * * * *", func() {
		log.Println("Executing flush cron job")
		flushAllDataFromHashSet()
	})

	cronScheduler.Start()

	select{}
}