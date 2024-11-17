package repository

import (
	"PencraftB/utils"
	"fmt"
	"log"
	"time"
	"github.com/joho/godotenv"
)


func ReadEnvFile() *map[string]string {
	envFile, err := godotenv.Read(".env")
	
	if err!=nil {
		log.Println(err)
		utils.Logger("Error loading .env File")
		log.Fatal("Fatal error!")
	} 

	// fmt.Println(envFile)

	return &envFile;
}	

type MongoDBConfig struct {
	URI string
	DatabaseName string
	ConnectionTimeout time.Duration
	PoolSize uint64
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
