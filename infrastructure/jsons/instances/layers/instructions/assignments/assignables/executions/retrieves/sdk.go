package retrieves

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
)

// NewAdapter creates a new adapter
func NewAdapter() retrieves.Adapter {
	builder := retrieves.NewBuilder()
	return createAdapter(
		builder,
	)
}