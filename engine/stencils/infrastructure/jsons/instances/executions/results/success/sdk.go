package success

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/success"
	json_outputs "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/executions/results/success/outputs"
	json_kinds "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/outputs/kinds"
)

// NewAdapter creates a new adapter
func NewAdapter() success.Adapter {
	outputAdapter := json_outputs.NewAdapter()
	kindAdapter := json_kinds.NewAdapter()
	builder := success.NewBuilder()
	return createAdapter(
		outputAdapter.(*json_outputs.Adapter),
		kindAdapter.(*json_kinds.Adapter),
		builder,
	)
}
