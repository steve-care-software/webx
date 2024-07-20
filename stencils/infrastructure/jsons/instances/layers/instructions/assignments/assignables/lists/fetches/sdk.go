package fetches

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/lists/fetches"
)

// NewAdapter creates a new adapter
func NewAdapter() fetches.Adapter {
	builder := fetches.NewBuilder()
	return createAdapter(
		builder,
	)
}
