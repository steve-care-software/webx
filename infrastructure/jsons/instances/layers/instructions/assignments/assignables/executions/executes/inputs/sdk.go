package inputs

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
)

// NewAdapter creates a new adapter
func NewAdapter() inputs.Adapter {
	builder := inputs.NewBuilder()
	return createAdapter(
		builder,
	)
}
