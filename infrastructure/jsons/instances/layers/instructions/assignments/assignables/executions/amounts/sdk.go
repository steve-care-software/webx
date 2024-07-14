package amounts

import "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/amounts"

// NewAdapter creates a new adapter
func NewAdapter() amounts.Adapter {
	builder := amounts.NewBuilder()
	return createAdapter(
		builder,
	)
}
