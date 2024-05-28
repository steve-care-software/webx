package interruptions

import (
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/interruptions"
	json_failures "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commands/results/interruptions/failures"
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
