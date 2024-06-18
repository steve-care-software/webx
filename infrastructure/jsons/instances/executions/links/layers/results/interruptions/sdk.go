package interruptions

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/interruptions"
	json_failures "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/links/layers/results/interruptions/failures"
)

// NewAdapter creates a new adapter
func NewAdapter() interruptions.Adapter {
	failureAdapter := json_failures.NewAdapter()
	builder := interruptions.NewBuilder()
	return createAdapter(
		failureAdapter.(*json_failures.Adapter),
		builder,
	)
}
