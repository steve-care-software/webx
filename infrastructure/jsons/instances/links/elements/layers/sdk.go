package layers

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
	json_instructions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions"
	json_output "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/outputs"
)

// NewAdapter creates a new adapter
func NewAdapter() layers.Adapter {
	instructionsAdapter := json_instructions.NewAdapter()
	outputAdapter := json_output.NewAdapter()
	layerBuilder := layers.NewLayerBuilder()
	builder := layers.NewBuilder()
	return createAdapter(
		instructionsAdapter.(*json_instructions.Adapter),
		outputAdapter.(*json_output.Adapter),
		layerBuilder,
		builder,
	)
}
