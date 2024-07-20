package executables

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/executables"
)

// NewAdapter creates a new adapter
func NewAdapter() executables.Adapter {
	builder := executables.NewBuilder()
	return createAdapter(
		builder,
	)
}
