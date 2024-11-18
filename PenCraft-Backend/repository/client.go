package repository

import (
	"PencraftB/models"
	relations "PencraftB/models/Relations"
	"PencraftB/utils"
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	client *mongo.Client
	config *MongoDBConfig
}

var (
	instance *DBClient
	once sync.Once
)


func NewDBClient() *DBClient {
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

	return instance;

}


func (db *DBClient) connect() {
	clientOptions := options.Client().
						ApplyURI(db.config.URI).
						SetConnectTimeout(db.config.ConnectionTimeout).
						SetMaxPoolSize(db.config.PoolSize)

	
	// establish connection
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("MongoDB Connection error : %v",err)
	}


	// Test the connection
	err = client.Ping(context.TODO(), nil)
	if err!= nil{
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	db.client = client;
	log.Println("Successfully connected to MongoDB")
}


func (db *DBClient) GetCollection(collectionName string ) *mongo.Collection {
	if db.client == nil {
		log.Fatal("MongoDB client is not initialized, to get collection Name")
	}

	return db.client.Database(db.config.DatabaseName).Collection(collectionName);
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


func (db *DBClient) SaveTagOnly(collectionName string, tag models.Tag) (interface{},error){
	
	var ctx, cancel = context.WithTimeout(context.Background(), 100* time.Second);
	defer cancel();

	collection := db.GetCollection(collectionName);
	resultChan := make( chan *mongo.InsertOneResult)
	errChan := make(chan error)

	var wg sync.WaitGroup;
	wg.Add(1)

	// goroutine
	go func ()  {
		defer wg.Done();

		result,err := collection.InsertOne(ctx, tag)
		if err != nil{
			log.Println("Could not save Tag in db !")
			errChan <- err;
			return;
		}

		resultChan <- result;
	}()


	// handling closing of channels
	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	select {
	case result := <- resultChan:
		return result,nil;
	case err := <-errChan:
		return nil, err;
	case <- ctx.Done():
		return nil, ctx.Err()
	}
}


func (db *DBClient) FetchAllTags()( interface{}, error){
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

	defer cancel();

	matchStage := bson.D{{Key: "$match",Value: bson.D{}}}

	groupStage := bson.D{{
		Key:"$group",
		Value: bson.D{
				{Key: "_id", Value: nil},
				{Key:"total_count", Value: bson.D{{Key: "$sum",Value: 1}}},
				{Key:"data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
			},
	}}

	projectStage := bson.D{
		{
			Key:"$project",
			Value:bson.D{
				{"_id",0},
				{"total_count",1},
				{"data",1},
			},
		},
	}

	collection := db.GetCollection(utils.ALL_TAG);
	result,err := collection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage,
	})

	if err!= nil{
		log.Println("Cannot fetch all tags")
		log.Println(err.Error())
		return nil,err;
	}


	var allTags []bson.M
	if err = result.All(ctx, &allTags); err != nil {
		log.Println(err.Error())
		log.Println("Failed while converting tags to bson.M[]");
		return nil,err;
	}
	
	
	return allTags,nil;
}


func (db *DBClient) SaveRelation(collectionName string, blog relations.R_Tag_Blog) (interface{}, error){
	var ctx, cancel = context.WithTimeout(context.Background(), 80 * time.Second)
	defer cancel();

	collection := db.GetCollection(collectionName)
	resultChan := make(chan *mongo.InsertOneResult)
	errChan := make(chan error)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		var blogToTag relations.R_Tag_Blog
		blogToTag.Blog_id = blog.Blog_id;
		blogToTag.Tag_id = blog.Tag_id;

		// checking for blog-relation-tag persistence
		result,err := collection.InsertOne(ctx, blogToTag)
		if err != nil{ 
			log.Println("Could not save blog in mongo-db (Second) !")
			errChan <- err;
			return;
		}
		
		resultChan <- result
	}()


	// close channels
	go func(){
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	select {
		case result := <- resultChan :
			return result,nil

		case err := <- errChan:
			return nil,err

		case <- ctx.Done():
			return nil, ctx.Err()
	}

}

func (db *DBClient) SaveBlog(collectionName string, blog models.Blog) (interface{}, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 80 * time.Second)
	defer cancel();

	collection := db.GetCollection(collectionName)

	resultChan := make(chan *mongo.InsertOneResult)
	errChan := make(chan error)

	var wg sync.WaitGroup;
	wg.Add(1)

	// go routine
	go func() {
		defer wg.Done()

		// checking for blog persistence
		result, err := collection.InsertOne(ctx, blog)
		if err != nil {
			log.Println("Could not save blog in mongo-db (First) !")
			errChan <- err;
			return;
		}

		resultChan <- result
	}()


	// close channels
	go func(){
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	select {
		case result := <- resultChan :
			return result,nil

		case err := <- errChan:
			return nil,err

		case <- ctx.Done():
			return nil, ctx.Err()
	}

}