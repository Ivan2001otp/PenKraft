package repository

import (
	"context"
	"log"
	"sync"
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
		log.Fatal("MongoDB client is not initialized, to get collection NAme")
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