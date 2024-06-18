package layers

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers"
	json_results "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/links/layers/results"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers"
)

// NewAdapter creates a new adapter
func NewAdapter() layers.Adapter {
	layerAdapter := json_layers.NewAdapter()
	resultAdapter := json_results.NewAdapter()
	builder := layers.NewBuilder()
	layerBuilder := layers.NewLayerBuilder()
	return createAdapter(
		layerAdapter.(*json_layers.Adapter),
		resultAdapter.(*json_results.Adapter),
		builder,
		layerBuilder,
	)
}
