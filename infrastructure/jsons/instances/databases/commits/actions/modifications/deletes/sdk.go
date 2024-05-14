package deletes

import (
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
)

// NewAdapter creates a new adapter
func NewAdapter() deletes.Adapter {
	builder := deletes.NewBuilder()
	return createAdapter(
		builder,
	)
}
