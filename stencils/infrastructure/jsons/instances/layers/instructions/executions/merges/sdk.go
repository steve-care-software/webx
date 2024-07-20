package merges

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/executions/merges"
)

// NewAdapter creates a new adapter
func NewAdapter() merges.Adapter {
	builder := merges.NewBuilder()
	return createAdapter(
		builder,
	)
}
