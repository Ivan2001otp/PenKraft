package repository

import (
	"PencraftB/models"
	"PencraftB/utils"
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	clientInstance *redis.Client
	redisOnce 			sync.Once
)

type RedisClient struct {
	client *redis.Client
}


func GetRedisInstance() *RedisClient {

	redisOnce.Do(func ()  {
		clientInstance = createRedisClient()
	})

	return &RedisClient{client: clientInstance}
}

func createRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB:0,
		MaxRetries: 2,
		DialTimeout: 5 * time.Second,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize: 10, 
		MinIdleConns: 4,
		IdleTimeout: 4* time.Minute,
		PoolTimeout: 4 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis")

	return rdb;
}

// Fetch the blogkey from message queue
func (r *RedisClient) PopBlogKeyFromQueue(ctx context.Context,queueName string)(*[]string,error){
	blogKey, err := r.client.BLPop(ctx, 0, queueName).Result()

	if err!= nil {
		return nil,err;
	}

	return &blogKey,nil;
} 


func (r *RedisClient) PushToMessageQueue(ctx context.Context, queueName string, redisKey string) error{
	err := r.client.LPush(ctx,queueName, redisKey).Err()

	return err;
}

// Fetch the blogData from cache by blogkey
func (r *RedisClient) PopBlogdataFromBlogkey(ctx context.Context, queueName string,blogKey string) ( *string , error){
	blogData, err := r.client.Get(ctx, blogKey).Result()

	if err != nil {
		log.Printf("Error retreiving the blogdata from blogKey from redisQueue. %v",err)
		return nil,err;
	}
	

	r.client.Del(ctx, blogKey)
	return &blogData, nil;
}


func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error{
	return r.client.Set(ctx, key, value, ttl).Err()
}


func (r *RedisClient) Get(ctx context.Context, key string)(string, error){
	return r.client.Get(ctx,key).Result()
}


func (r *RedisClient) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}


func (r *RedisClient) Close() error {
	return r.client.Close()
}


// fetches all blogs from Redis using HGetAll
func (r *RedisClient) FetchAllBlogfromRedis(ctx context.Context) ([]models.Blog, error) {

	// fetch all blogs from the "blogs" hash
	result,err := r.client.HGetAll(ctx, utils.REDIS_BLOG_COLLECTION).Result()
	if err != nil {
		return nil, err;
	}

	var listOfBlog []models.Blog;
	for _, blogData := range result {
		var blog models.Blog

		// Unmarshal each student's JSON data
		err := json.Unmarshal([]byte(blogData), &blog)
		if err!= nil {
			return nil, err;
		}

		listOfBlog = append(listOfBlog, blog)
	}

	return listOfBlog, nil;
}


// saving blogs to redis using HSet
func (r *RedisClient) SaveAllBlogtoRedis(ctx context.Context, listOfBlog []models.Blog) {

	for _,blogItem := range listOfBlog {
		
		blogData, err := json.Marshal(blogItem)
		if err != nil {
			log.Fatalf("Error marshalling student: %v", err)
		}

		// store each blog in Redis Hash with BlogId as key.
		err = r.client.HSet(ctx, utils.REDIS_BLOG_COLLECTION, blogData).Err()
		if err != nil {
			log.Fatalf("Error storing blog in Redis: %v", err)
		}

	}
}