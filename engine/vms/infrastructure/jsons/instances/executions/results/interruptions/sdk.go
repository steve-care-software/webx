package interruptions

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/interruptions"
	json_failures "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/executions/results/interruptions/failures"
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
