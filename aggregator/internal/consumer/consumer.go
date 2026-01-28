package consumer

import (
	"log"

	"github.com/IBM/sarama"
)

func NewKafkaConsumer(broker string, topics string) *sarama.PartitionConsumer {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	consumer, err := sarama.NewConsumer([]string{broker}, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topics, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		log.Printf("Received: %s\n", string(msg.Value))
	}
	return &partitionConsumer
}
