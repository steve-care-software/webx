package layers

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers"
	json_instructions "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/instructions"
	json_output "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/outputs"
	json_references "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/layers/references"
)

// NewAdapter creates a new adapter
func NewAdapter() layers.Adapter {
	instructionsAdapter := json_instructions.NewAdapter()
	outputAdapter := json_output.NewAdapter()
	referenceAdapter := json_references.NewAdapter()
	builder := layers.NewBuilder()
	return createAdapter(
		instructionsAdapter.(*json_instructions.Adapter),
		outputAdapter.(*json_output.Adapter),
		referenceAdapter.(*json_references.Adapter),
		builder,
	)
}
