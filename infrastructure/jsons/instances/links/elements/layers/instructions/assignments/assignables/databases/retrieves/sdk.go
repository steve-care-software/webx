package retrieves

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/databases/retrieves"
)

// NewAdapter creates a new adapter
func NewAdapter() retrieves.Adapter {
	builder := retrieves.NewBuilder()
	return createAdapter(
		builder,
	)
}
