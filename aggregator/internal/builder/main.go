package builder

import (
	"github.com/IBM/sarama"
	"github.com/gabriwl165/enterprise-integrations-pattern/aggregator/internal/core"
	"github.com/gabriwl165/enterprise-integrations-pattern/aggregator/internal/usecases"
)

func NewAggregator(producer sarama.SyncProducer) *core.Aggregator {
	return &core.Aggregator{
		Aggregates:             make(map[string][]core.Aggregate),
		PublishEventHandler:    &usecases.PublishEventHandler{Producer: producer},
		IsCompleteEventHandler: &usecases.IsCompleteEventHandler{},
	}
}
