package heads

import "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/heads"

// NewAdapter creates a new adapter
func NewAdapter() heads.Adapter {
	builder := heads.NewBuilder()
	return createAdapter(
		builder,
	)
}
