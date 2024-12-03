package repository

import (
	"context"
	"fmt"
	"log"
	"time"
	"sync"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBConfig struct {
	URI string
	DatabaseName string
	ConnectionTimeout time.Duration
	PoolSize uint64
}

// this is the changestream provider that manages stream for a given collection.
type ChangeStreamManager struct {
	client *mongo.Client
	collection *mongo.Collection
	changeStream *mongo.ChangeStream
	streamCtx  context.Context
	exCancelFunc context.CancelFunc
	once sync.Once
	mu	sync.Mutex
}

// creates and return the singleton object of ChangeStreamManager.
func NewChangeStreamManager( client *mongo.Client,collection string) *ChangeStreamManager {
	envFile := make(map[string]string)
	envFile = *ReadEnvFile()

	databaseName := envFile["DATABASE_NAME"];
	if databaseName=="" {
		databaseName = "PENKRAFT"
	}

	return &ChangeStreamManager{
		client: client,
		collection: client.Database(databaseName).Collection(collection),
	}
}


func ReadEnvFile() *map[string]string {
	envFile, err := godotenv.Read("../.env")
	
	if err!=nil {
		log.Println(err)
		log.Println("Error loading .env File")
		log.Fatal("Fatal error!")
	} 

	// fmt.Println(envFile)

	return &envFile;
}	


func NewMongoDBConfig() (*MongoDBConfig, error) {

	envFile := make(map[string]string)

	envFile = *ReadEnvFile();
	uri := envFile["MONGO_URL"]

	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	databaseName := envFile["DATABASE_NAME"];
	if databaseName=="" {
		databaseName = "PENKRAFT"
	}

	// Connection timeout, default to 10 seconds
	connectionTimeout := 10 * time.Second;

	if timeoutStr := envFile["MONGO_TIMEOUT"]; timeoutStr!= "" {
		timeout, err := time.ParseDuration(timeoutStr)

		if err!=nil {
			log.Println("Invalid MONGO_TIMEOUT ,using default: %v", err)
		} else {
			connectionTimeout = timeout
		}
	}


	// pool size,default to 10
	poolSize := uint64(10)
	if poolSizeStr := envFile["MONGO_POOL_SIZE"]; poolSizeStr!= "" {
		fmt.Sscanf(poolSizeStr,"%d",&poolSize)
	}


	return &MongoDBConfig{
		URI: uri,
		DatabaseName: databaseName,
		ConnectionTimeout: connectionTimeout,
		PoolSize: poolSize,
	},nil

}
