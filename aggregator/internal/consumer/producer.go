package consumer

import (
	"log"

	"github.com/IBM/sarama"
)

func NewKafkaProducer(broker string) sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer([]string{broker}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return producer
}
