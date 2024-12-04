package main

import (
	Util "PencraftB/utils"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/IBM/sarama"
	"github.com/elastic/go-elasticsearch/v7"
  "github.com/elastic/go-elasticsearch/v7/esapi"
)

var kafkaBroker  = Util.KAFKA_BROKER //broker address
var kafkaTopi = Util.KAFKA_TOPIC //kafka topic to consume mongochanges
var esClient *elasticsearch.Client

// initialize elasticsearch client
func init() {
	
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Address: []string{"http://"+Util.KAFKA_BROKER},
	})

	if err != nil {
		log.Fatalf("Error on creating elastic search client : %v" ,err)
	}
}


func ProcessEventToElasticSearch(event map[string]interface{}) {
	// convert event to json.
	eventData, err := json.Marshal(event)

	if err != nil {
		log.Printf("(ProcessEventToElasticSearch)Error marshalling event : %v",err)
		return;
	}


	// index the document
	req := esapi.IndexRequest{
		Index: "mongo-events",
		DocumentID: "",
		Body: bytes.NewReader(eventData),
		Refresh: "true",
	}

	// send request to elasticsearch
	resp,err := req.Do(context.Background(), esClient)
	if err != nil {
		log.Printf("Error indexing event to ElasticSearch: %v", err)
		return;
	}

	defer resp.Body.Close()


	if resp.IsError() {
		log.Printf("Error indexing event: %v", resp);
	} else {
		log.Printf("Event indexed to elasticsearch successfully")
	}

}


func ConsumeKafkaMessages() {
	// kafka consumer setup
	consumer, err := sarama.NewConsumer([] string{Util.KAFKA_BROKER}, nil)
	if err != nil {
		log.Fatalf("Error creating kafka consumer : %v" ,err)
	}

	defer consumer.Close()

	//setting up channels to capture system signals on interrupt
	stopChan := make(chan os.Signal,1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)


	//subscribe to kafka topic
	partitions, err := consumer.Partitions(Util.KAFKA_TOPIC)
	if err != nil {
		log.Fatalf("Error fetching partitions. %v",err)
		return;
	}

	for _, partition := range partitions {
		pc,err := consumer.ConsumePartition(Util.KAFKA_TOPIC, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Error consuming Partition: %d, %v",partition,err)
			return;
		}
		defer pc.Close()

		go func() {
			for {
				select {	
				case msg:= <-pc.Messages() :
					// process the kafka event
					log.Printf("Processing kafka messages")
					log.Printf("Recieved message is %s",string(msg.Value))

					//assuming data in JSON
					var event map[string]interface{}
					err := json.Unmarshal(msg.Value, &event)
					if err != nil {
						log.Printf("Error unmarshalling message : %v", err)
						continue;
					}

					ProcessEventToElasticSearch(event)
					

				case <-stopChan:
					log.Println("Gracefully shutting down Kafka consumer...")
					return

				}
			}
		}()
	}

	<-stopChan
	log.Println("Exiting consumer...");
}


func main() {
	ConsumeKafkaMessages()
}