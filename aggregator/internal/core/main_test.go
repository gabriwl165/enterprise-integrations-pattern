package core

import (
	"testing"
)

type TestAggregatorPublishEventHandler struct{}

func (t *TestAggregatorPublishEventHandler) Publish(aggregates []Aggregate) error {
	return nil
}

type TestAggregatorIsCompleteEventHandler struct{}

func (t *TestAggregatorIsCompleteEventHandler) IsComplete(messages []Aggregate) bool {
	return len(messages) == 5
}

func TestAggregator_AddMessage(t *testing.T) {
	aggregator := Aggregator{
		Aggregates:             make(map[string][]Aggregate),
		PublishEventHandler:    &TestAggregatorPublishEventHandler{},
		IsCompleteEventHandler: &TestAggregatorIsCompleteEventHandler{},
	}
	aggregator.AddMessage(map[string]interface{}{"id": "1", "message": "Hello, World!"})
	aggregates, _ := aggregator.GetCorrelationId("1")
	if len(aggregates) != 1 {
		t.Errorf("expected aggregator to have id '1' with length 1, got: %v", aggregates)
	}
}

func TestAggregator_IsComplete(t *testing.T) {
	aggregator := Aggregator{
		Aggregates:             make(map[string][]Aggregate),
		PublishEventHandler:    &TestAggregatorPublishEventHandler{},
		IsCompleteEventHandler: &TestAggregatorIsCompleteEventHandler{},
	}
	aggregator.AddMessage(map[string]interface{}{"id": "1", "message": "Hello, World!"})
	aggregates, _ := aggregator.GetCorrelationId("1")
	if aggregator.IsCompleteEventHandler.IsComplete(aggregates) {
		t.Errorf("expected aggregator not to be complete")
	}
}
