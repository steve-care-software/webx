package failures

import (
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/interruptions/failures"
)

// NewAdapter creates a new adapter
func NewAdapter() failures.Adapter {
	builder := failures.NewBuilder()
	return createAdapter(
		builder,
	)
}
