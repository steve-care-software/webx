package deletes

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/deletes"
)

// NewAdapter creates a new adapter
func NewAdapter() deletes.Adapter {
	builder := deletes.NewBuilder()
	return createAdapter(
		builder,
	)
}
