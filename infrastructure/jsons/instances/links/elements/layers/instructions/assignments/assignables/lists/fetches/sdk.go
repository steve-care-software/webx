package fetches

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/lists/fetches"
)

// NewAdapter creates a new adapter
func NewAdapter() fetches.Adapter {
	builder := fetches.NewBuilder()
	return createAdapter(
		builder,
	)
}
