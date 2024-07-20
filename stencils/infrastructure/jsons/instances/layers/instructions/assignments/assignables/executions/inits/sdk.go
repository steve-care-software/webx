package inits

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/executions/inits"
)

// NewAdapter creates a new adapter
func NewAdapter() inits.Adapter {
	builder := inits.NewBuilder()
	return createAdapter(
		builder,
	)
}
