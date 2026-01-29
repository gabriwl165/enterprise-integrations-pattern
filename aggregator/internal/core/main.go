package core

import (
	"fmt"
	"time"
)

type Aggregate struct {
	CorrelationID string
	Messages      []interface{}
	CreatedOn     time.Time
}

type AggregatorPublishEventHandler interface {
	Publish(aggregates []Aggregate) error
}

type AggregatorIsCompleteEventHandler interface {
	IsComplete(messages []Aggregate) bool
}

type Aggregator struct {
	Aggregates             map[string][]Aggregate
	PublishEventHandler    AggregatorPublishEventHandler
	IsCompleteEventHandler AggregatorIsCompleteEventHandler
}

func (a *Aggregator) AddMessage(msg map[string]interface{}) error {
	id, ok := msg["id"].(string)
	if !ok {
		return fmt.Errorf("id is required and must be a string")
	}

	a.AddAggregate(id, Aggregate{
		CorrelationID: id,
		Messages:      []interface{}{msg},
		CreatedOn:     time.Now(),
	})

	messages, _ := a.GetCorrelationId(id)
	if a.IsCompleteEventHandler.IsComplete(messages) {
		a.PublishEventHandler.Publish(messages)
	}

	return nil
}

func (a *Aggregator) GetCorrelationId(correlationID string) ([]Aggregate, error) {
	aggregates, ok := a.Aggregates[correlationID]
	if !ok {
		return nil, fmt.Errorf("aggregate not found")
	}

	return aggregates, nil
}

func (a *Aggregator) AddAggregate(id string, aggregate Aggregate) error {
	a.Aggregates[id] = append(a.Aggregates[id], aggregate)
	return nil
}
