package config

import (
	"PencraftB/utils"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

var (
	kafkaProducer sarama.SyncProducer
	once2 sync.Once
)

// single ton instance of Kafka
func GetKafkaProducer(brokerList []string) sarama.SyncProducer{ 
	once2.Do(func() {
		log.Println("Initializing Kafka Producer....")

		config := sarama.NewConfig()

		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Retry.Max = utils.NUMBER_OF_RETRIES
		config.Producer.Return.Successes = true;


		var err error;
		tempKafkaProducer, err := sarama.NewSyncProducer(brokerList, config);
		if err != nil {
			log.Fatalf("Error initializing kafka producer: %v" ,err);
			return;
		}
		kafkaProducer = tempKafkaProducer;
		log.Println("Successfully instantiated kafka producer !")
	})

	return kafkaProducer;
}