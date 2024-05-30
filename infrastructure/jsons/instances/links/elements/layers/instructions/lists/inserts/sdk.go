package inserts

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/lists/inserts"
)

// NewAdapter creates a new adapter
func NewAdapter() inserts.Adapter {
	builder := inserts.NewBuilder()
	return createAdapter(
		builder,
	)
}
