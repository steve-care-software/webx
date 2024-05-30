package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/databases/databases"
)

// NewAdapter creates a new adapter
func NewAdapter() databases.Adapter {
	builder := databases.NewBuilder()
	return createAdapter(
		builder,
	)
}
