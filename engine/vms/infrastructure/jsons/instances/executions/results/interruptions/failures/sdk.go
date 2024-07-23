package failures

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/interruptions/failures"
)

// NewAdapter creates a new adapter
func NewAdapter() failures.Adapter {
	builder := failures.NewBuilder()
	return createAdapter(
		builder,
	)
}
