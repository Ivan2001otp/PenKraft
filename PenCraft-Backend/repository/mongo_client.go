package repository

import (
	"PencraftB/config"
	"PencraftB/models"
	relations "PencraftB/models/Relations"
	"PencraftB/utils"
	"context"
	"encoding/json"
	"fmt"

	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"go.mongodb.org/mongo-driver/bson"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	client *mongo.Client
	config *MongoDBConfig
}

var (
	instance *DBClient
	once     sync.Once
)

// **********************************
// mongodb ChangeStream operations
// **********************************

// GetMongoDbObserver ensures there exists a single active change-stream per collection. 
func (observer *ChangeStreamManager) GetMongoDbObserver() (*mongo.ChangeStream, error) {
	observer.mu.Lock()
	defer observer.mu.Unlock()

	observer.once.Do(func() {
		//setting up context with cancellation
		observer.streamCtx, observer.exCancelFunc = context.WithCancel(context.Background())

		// start watching the collection
		changeStream, err := observer.collection.Watch(observer.streamCtx, mongo.Pipeline{})
		if err != nil {
			log.Printf("Error starting change stream : %v", err)
			return
		}

		observer.changeStream = changeStream
		log.Println("ChangeStream started for collection %s", observer.collection.Name())
	})

	if observer.changeStream == nil {
		return nil, fmt.Errorf("change stream initialization faild for collection %s", observer.collection.Name())
	}

	return observer.changeStream, nil
}

// CLoseMongoDbObserver closes the given single active changestream
func (observer *ChangeStreamManager) CloseMongoDbObserver() error {
	observer.mu.Lock()
	defer observer.mu.Unlock();

	if observer.changeStream != nil {
		err := observer.changeStream.Close(observer.streamCtx)

		if err != nil {
			log.Printf("Error occured while closing the change-Stream : %v", err)
			return err;
		}
		log.Printf("ChangeStream closed for collection %s",observer.collection.Name())
	}


	// reset the changeStream attributes
	observer.changeStream = nil;
	observer.once = sync.Once{}
	if observer.exCancelFunc != nil {
		observer.exCancelFunc()
	}

	return nil;
}


func (observer *ChangeStreamManager) MonitorChanges() {
	watchStream, err := observer.GetMongoDbObserver();

	if err != nil {
		log.Fatal("Failed to get ChangeStream.",err)
		return;
	}

	// Process the onchange events to feed the data to Elastic via kafka
	defer watchStream.Close(context.Background())
	brokerList := []string{"localhost:9092"}
	kafkaProducer := config.GetKafkaProducer(brokerList)

	for watchStream.Next(context.Background()) {
		var event bson.M;
		if err := watchStream.Decode(&event); err != nil {
			log.Println("Error decoding event : %v" ,err);
			continue;
		}

		
		// serialize the event to json,in order to send Kafka
		
		message,err := json.Marshal(event)
		if err != nil {
			log.Println("(monmgo_client.go -> MonitorChanges)Marshalling Failed");
			continue;
		}

		msg:= &sarama.ProducerMessage{
			Topic: utils.KAFKA_TOPIC,
			Value: sarama.StringEncoder(message),
		}

		_,_,err = kafkaProducer.SendMessage(msg)

		if err != nil {
			log.Println("Error sending message to Kafka : ",err)
		} else {
			log.Println("Message sent successfully !")
		}

	}

	log.Println("Stopped watching changes.")
}

//************************************************************************************
//************************************************************************************

func GetMongoDBClient() *DBClient {
	once.Do(func() {

		config, err := NewMongoDBConfig()
		// Load configuration
		if err != nil {
			log.Fatalf("Error loading mongoDB configuration: %v", err)
		}

		instance = &DBClient{
			config: config,
		}

		instance.connect()
	})

	return instance
}

func (db *DBClient) connect() {
	clientOptions := options.Client().
		ApplyURI(db.config.URI).
		SetConnectTimeout(db.config.ConnectionTimeout).
		SetMaxPoolSize(db.config.PoolSize)

	// establish connection
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("MongoDB Connection error : %v", err)
	}

	// Test the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	db.client = client
	log.Println("Successfully connected to MongoDB")
}

func (db *DBClient) GetCollection(collectionName string) *mongo.Collection {
	if db.client == nil {
		log.Fatal("MongoDB client is not initialized, to get collection Name")
	}

	return db.client.Database(db.config.DatabaseName).Collection(collectionName)
}

func (db *DBClient) Close() {
	if db.client != nil {
		err := db.client.Disconnect(context.TODO())

		if err != nil {
			log.Printf("Error closing mongoDB connection: %v", err)
		} else {
			log.Println("MongoDB connection closed")
		}
	}
}

/*
************************************************************************************
TAG OPERATIONS
************************************************************************************
*/
func (db *DBClient) DeleteAllTags(ctx context.Context) error {

	collection := db.GetCollection(utils.ALL_TAG)

	_, err := collection.DeleteMany(ctx, bson.M{})

	if err != nil {
		log.Println("Could not delete  all tags")
		return err
	}

	return nil
}

func (db *DBClient) SoftDeleteTagbyId(ctx context.Context, tagId string) error {

	// mention softdelete param and set it to true.
	updatedBody := bson.M{
		"$set": bson.M{
			"is_delete": true,
		},
	}

	var upsert bool = true

	filter := bson.M{"tag_id": tagId}
	option := options.UpdateOptions{
		Upsert: &upsert,
	}

	collection := db.GetCollection(utils.ALL_TAG)
	result, err := collection.UpdateOne(ctx, filter, updatedBody, &option)

	if err != nil {
		log.Println("Could not soft delete the tag")
		log.Printf("%v", err)
		return err
	}

	log.Printf("Update count is : %v", result.UpsertedCount)
	return nil
}

// tag handlers
func (db *DBClient) SaveTagOnly(ctx context.Context, tag models.Tag) (interface{}, error) {

	collection := db.GetCollection(utils.ALL_TAG)
	resultChan := make(chan *mongo.InsertOneResult)
	errChan := make(chan error)

	var wg sync.WaitGroup
	wg.Add(1)

	// goroutine
	go func() {
		defer wg.Done()

		result, err := collection.InsertOne(ctx, tag)
		if err != nil {
			log.Println("Could not save Tag in db !")
			errChan <- err
			return
		}

		resultChan <- result
	}()

	// handling closing of channels
	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	select {
	case result := <-resultChan:
		return result, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (db *DBClient) FetchAllTags(ctx context.Context) (interface{}, error) {

	matchStage := bson.D{{Key: "$match", Value: bson.D{}}}

	// groupStage := bson.D{{
	// 	Key:"$group",
	// 	Value: bson.D{
	// 			{Key: "_id", Value: nil},
	// 			{Key:"total_count", Value: bson.D{{Key: "$sum",Value: 1}}},
	// 			{Key:"data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
	// 		},
	// }}

	// projectStage := bson.D{
	// 	{
	// 		Key:"$project",
	// 		Value:bson.D{
	// 			{"_id",0},
	// 			{"total_count",1},
	// 			{"data",1},
	// 		},
	// 	},
	// }

	collection := db.GetCollection(utils.ALL_TAG)
	result, err := collection.Aggregate(ctx, mongo.Pipeline{
		matchStage,
	})

	if err != nil {
		log.Println("Cannot fetch all tags")
		log.Println(err.Error())
		return nil, err
	}

	var allTags []bson.M
	if err = result.All(ctx, &allTags); err != nil {
		log.Println(err.Error())
		log.Println("Failed while converting tags to bson.M[]")
		return nil, err
	}

	return allTags, nil
}

// Fetch tag by tagid
func (db *DBClient) FetchTagbyId(ctx context.Context, tagId string) (*models.Tag, error) {

	collection := db.GetCollection(utils.ALL_TAG)

	filter := bson.M{"tag_id": tagId}

	result := collection.FindOne(ctx, filter)

	var tag models.Tag
	err := result.Decode(&tag)
	if err != nil {
		log.Println("Could not decode the tag")
		return nil, err
	}

	return &tag, nil
}

// Delete tag by id
func (db *DBClient) HardDeleteTagbyId(ctx context.Context, tagId string) error {

	filter := bson.M{"tag_id": tagId}
	collection := db.GetCollection(utils.ALL_TAG)

	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Printf("Error while deleting tag by id %v", err)
		return err
	}

	return nil
}

/*
************************************************************************************
BLOG OPERATIONS
************************************************************************************
*/

// Blog handlers
func (db *DBClient) SaveBlog(blog models.Blog) (interface{}, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 80*time.Second)
	defer cancel()

	collection := db.GetCollection(utils.BLOG_COLLECTION)

	resultChan := make(chan *mongo.InsertOneResult)
	errChan := make(chan error)

	var wg sync.WaitGroup
	wg.Add(1)

	// go routine
	go func() {
		defer wg.Done()

		// checking for blog persistence
		result, err := collection.InsertOne(ctx, blog)
		if err != nil {
			log.Println("Could not save blog in mongo-db (First) !")
			errChan <- err
			return
		}

		resultChan <- result
	}()

	// close channels
	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	select {
	case result := <-resultChan:
		return result, nil

	case err := <-errChan:
		return nil, err

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (db *DBClient) FetchAllBlogs(ctx context.Context) ([]models.Blog, error) {

	matchStage := bson.D{{Key: "$match", Value: bson.D{}}}

	// groupStage := bson.D{{
	// 	Key: "$group",
	// 	Value: bson.D{
	// 		{Key:"_id", Value:nil},
	// 		{Key:"total_count",Value:bson.D{{Key: "$sum",Value: 1}}},
	// 		{Key: "data", Value: bson.D{{Key: "$push",Value: "$$ROOT"}}},
	// 	},
	// }}

	// projectStage := bson.D{
	// 	{
	// 		Key: "$project",
	// 		Value: bson.D{
	// 			{"_id",0},
	// 			{"total_count",1},
	// 			{"data",1},
	// 		},
	// 	},
	// }

	collection := db.GetCollection(utils.BLOG_COLLECTION)
	cursorResult, err := collection.Aggregate(ctx, mongo.Pipeline{
		matchStage,
	})

	if err != nil {
		log.Println("Cannot fetch all blogs at client.go !")
		log.Println(err.Error())
		return nil, err
	}

	var listOfBlog []models.Blog
	for cursorResult.Next(ctx) {
		var blogPost models.Blog

		if err := cursorResult.Decode(&blogPost); err != nil {
			log.Println("failed to decode blog one by one !")
			return nil, err
		}
		listOfBlog = append(listOfBlog, blogPost)
	}

	if err := cursorResult.Err(); err != nil {
		log.Println("there was error in cursor in mongo")
		return nil, err
	}

	return listOfBlog, nil
}

func (db *DBClient) UpdateBlog(ctx context.Context, blog models.Blog) error {

	updatedBody := bson.M{
		"$set": bson.M{
			"title":      blog.Title,
			"excerpt":    blog.Excerpt,
			"tag_id":     blog.Tag_id,
			"updated_at": blog.Updated_at,
			"body":       blog.Body,
			"image":      blog.Image,
			"slug":       blog.Slug,
		},
	}

	collection := db.GetCollection(utils.BLOG_COLLECTION)
	filter := bson.M{"blog_id": blog.Blog_id}
	var upsert bool = false

	option := options.UpdateOptions{
		Upsert: &upsert,
	}
	result, err := collection.UpdateOne(ctx, filter, updatedBody, &option)

	if err != nil {
		log.Printf("Failed to update blog - %v", err)
		return err
	}

	if result.MatchedCount == 0 {
		log.Println("No matched count")
		return fmt.Errorf("Blog %s does not exist", blog.Blog_id)
	}

	log.Printf("updated %s blog .", blog.Blog_id)
	return nil
}

func (db *DBClient) FetchBlogbyBlogId(ctx context.Context, blogId string) (*models.Blog, error) {

	collection := db.GetCollection(utils.BLOG_COLLECTION)

	var blog models.Blog

	filter := bson.M{"blog_id": blogId}
	err := collection.FindOne(ctx, filter).Decode(&blog)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Blog with %s does not exist in mongo.", blogId)
			return nil, err
		}
		log.Println("Failed to find specific blog : ", err.Error())
		return nil, err
	}

	return &blog, nil
}

func (db *DBClient) DeleteAllBlogs(ctx context.Context) error {

	filter := bson.M{} // empty filter matches
	collection := db.GetCollection(utils.BLOG_COLLECTION)
	result, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		return fmt.Errorf("could not delete records %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no records found to delete")
	}

	log.Printf("Deleted %v records.", result.DeletedCount)
	return nil
}

func (db *DBClient) SoftDeleteBlogbyId(ctx context.Context, blogId string) error {
	collection := db.GetCollection(utils.BLOG_COLLECTION)

	filter := bson.M{"blog_id": blogId}
	updateBody := bson.M{
		"$set": bson.M{
			"is_delete": true,
		},
	}

	var upsert bool = true
	option := options.UpdateOptions{
		Upsert: &upsert,
	}

	result, err := collection.UpdateOne(ctx, filter, updateBody, &option)

	if err != nil {
		log.Printf("Could not delete blog(DeleteBlogbyId): %v", err)
		return fmt.Errorf("could not delete blog: %v", err)

	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no blog found with ID : %s", blogId)
	}

	return nil
}

// **********************************************************************
// ***********************************************************************
func (db *DBClient) DeleteAllRelations(collectionName string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 80*time.Second)
	defer cancel()
	collection := db.GetCollection(collectionName)

	filter := bson.M{} // delete all records.
	result, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		log.Fatalf("Could not delete all relations - %v", err)
		return err
	}
	log.Println("Deleted relations - ", result.DeletedCount)

	return nil
}

// relation handlers
func (db *DBClient) SaveRelation(collectionName string, blog relations.R_Tag_Blog) (interface{}, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 80*time.Second)
	defer cancel()
	collection := db.GetCollection(collectionName)
	resultChan := make(chan *mongo.InsertOneResult)
	errChan := make(chan error)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		var blogToTag relations.R_Tag_Blog
		blogToTag.Blog_id = blog.Blog_id
		blogToTag.Tag_id = blog.Tag_id

		// checking for blog-relation-tag persistence
		result, err := collection.InsertOne(ctx, blogToTag)
		if err != nil {
			log.Println("Could not save blog in mongo-db (Second) !")
			errChan <- err
			return
		}

		resultChan <- result
	}()

	// close channels
	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	select {
	case result := <-resultChan:
		return result, nil

	case err := <-errChan:
		return nil, err

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
