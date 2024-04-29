package inserts

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/inserts"
)

// NewAdapter creates a new adapter
func NewAdapter() inserts.Adapter {
	builder := inserts.NewBuilder()
	return createAdapter(
		builder,
	)
}
