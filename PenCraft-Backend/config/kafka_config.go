package config

import (
	"PencraftB/utils"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

var (
	kafkaProducer sarama.SyncProducer
	once2         sync.Once
)

var (
	consumerConfigObj *sarama.Config
	consumerConfigOnce sync.Once
)

// single ton instance of Kafka
func GetKafkaProducer(brokerList []string) sarama.SyncProducer {
	once2.Do(func() {
		log.Println("Initializing Kafka Producer....")

		config := FetchConsumerConfig()

		var err error
		tempKafkaProducer, err := sarama.NewSyncProducer(brokerList, config)
		if err != nil {
			log.Fatalf("Error initializing kafka producer: %v", err)
			return
		}
		kafkaProducer = tempKafkaProducer
		log.Println("Successfully instantiated kafka producer !")
	})

	return kafkaProducer
}



func FetchConsumerConfig() *sarama.Config{

	consumerConfigOnce.Do(func() {
		log.Println("Creating consumer config !")
		consumerConfigObj = sarama.NewConfig();
		consumerConfigObj.Producer.RequiredAcks = sarama.WaitForAll;
		consumerConfigObj.Producer.Retry.Max = utils.NUMBER_OF_RETRIES;
		consumerConfigObj.Producer.Return.Successes = true;
	})
	
	log.Println("Successfully created consumer config !")
	return consumerConfigObj;
}