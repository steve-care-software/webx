package executions

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
	json_results "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/executions/results"
	json_layers "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers"
)

// NewAdapter creates a new adapter
func NewAdapter() executions.Adapter {
	resultAdapter := json_results.NewAdapter()
	layerAdapter := json_layers.NewAdapter()
	builder := executions.NewBuilder()
	executionsBuilder := executions.NewExecutionBuilder()
	return createAdapter(
		resultAdapter.(*json_results.Adapter),
		layerAdapter.(*json_layers.Adapter),
		builder,
		executionsBuilder,
	)
}
