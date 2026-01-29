package usecases

import "github.com/gabriwl165/enterprise-integrations-pattern/aggregator/internal/core"

type IsCompleteEventHandler struct {
}

func (i *IsCompleteEventHandler) IsComplete(messages []core.Aggregate) bool {
	return len(messages) == 5
}
