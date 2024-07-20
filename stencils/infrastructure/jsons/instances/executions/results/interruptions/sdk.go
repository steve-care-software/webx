package interruptions

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/executions/results/interruptions"
	json_failures "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/executions/results/interruptions/failures"
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
