package fetches

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/lists/fetches"
)

// NewAdapter creates a new adapter
func NewAdapter() fetches.Adapter {
	builder := fetches.NewBuilder()
	return createAdapter(
		builder,
	)
}
