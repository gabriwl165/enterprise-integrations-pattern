package usecases

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/gabriwl165/enterprise-integrations-pattern/aggregator/internal/core"
)

type PublishEventHandler struct {
	Producer sarama.SyncProducer
}

func (p *PublishEventHandler) Publish(aggregates []core.Aggregate) error {
	var messages []interface{}
	for _, aggregate := range aggregates {
		messages = append(messages, aggregate.Messages...)
	}

	msgBytes := []byte(fmt.Sprintf("%v", messages))

	kafkaMsg := &sarama.ProducerMessage{
		Topic: "aggregated-topic",
		Value: sarama.ByteEncoder(msgBytes),
	}

	_, _, err := p.Producer.SendMessage(kafkaMsg)
	if err != nil {
		return err
	}
	return nil
}

func NewPublishEventHandler(producer sarama.SyncProducer) *PublishEventHandler {
	return &PublishEventHandler{Producer: producer}
}
