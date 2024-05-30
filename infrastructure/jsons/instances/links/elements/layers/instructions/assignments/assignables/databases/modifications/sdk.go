package modifications

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/databases/modifications"
)

// NewAdapter creates a new adapter
func NewAdapter() modifications.Adapter {
	builder := modifications.NewBuilder()
	return createAdapter(
		builder,
	)
}
