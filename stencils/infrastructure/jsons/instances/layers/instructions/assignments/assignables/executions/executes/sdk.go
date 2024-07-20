package executes

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	json_inputs "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
)

// NewAdapter creates a new adapter
func NewAdapter() executes.Adapter {
	inputAdapter := json_inputs.NewAdapter()
	builder := executes.NewBuilder()
	return createAdapter(
		inputAdapter.(*json_inputs.Adapter),
		builder,
	)
}
