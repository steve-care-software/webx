package deletes

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/deletes"
)

// NewAdapter creates a new adapter
func NewAdapter() deletes.Adapter {
	builder := deletes.NewBuilder()
	return createAdapter(
		builder,
	)
}
