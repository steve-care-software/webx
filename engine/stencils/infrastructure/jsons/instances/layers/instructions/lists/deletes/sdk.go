package deletes

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/lists/deletes"
)

// NewAdapter creates a new adapter
func NewAdapter() deletes.Adapter {
	builder := deletes.NewBuilder()
	return createAdapter(
		builder,
	)
}
