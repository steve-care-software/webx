package commits

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/commits"
)

// NewAdapter creates a new adapter
func NewAdapter() commits.Adapter {
	builder := commits.NewBuilder()
	return createAdapter(
		builder,
	)
}
