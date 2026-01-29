package main

import (
	"encoding/json"

	"github.com/gabriwl165/enterprise-integrations-pattern/aggregator/internal/builder"
	"github.com/gabriwl165/enterprise-integrations-pattern/aggregator/internal/consumer"
)

func main() {
	kafkaProducer := consumer.NewKafkaProducer("localhost:9092")
	aggregator := builder.NewAggregator(kafkaProducer)
	kafkaConsumer := consumer.NewKafkaConsumer("localhost:9092", "my-topic")
	for msg := range kafkaConsumer.Messages() {
		var payload map[string]interface{}
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			continue
		}
		aggregator.AddMessage(payload)
	}
}
