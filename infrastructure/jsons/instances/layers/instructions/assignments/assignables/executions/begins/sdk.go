package begins

import "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/begins"

// NewAdapter creates a new adapter
func NewAdapter() begins.Adapter {
	builder := begins.NewBuilder()
	return createAdapter(
		builder,
	)
}
