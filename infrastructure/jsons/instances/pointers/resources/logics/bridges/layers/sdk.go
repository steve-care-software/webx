package layers

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
	json_instructions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions"
	json_output "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/outputs"
)

// NewAdapter creates a new adapter
func NewAdapter() layers.Adapter {
	instructionsAdapter := json_instructions.NewAdapter()
	outputAdapter := json_output.NewAdapter()
	builder := layers.NewBuilder()
	return createAdapter(
		instructionsAdapter.(*json_instructions.Adapter),
		outputAdapter.(*json_output.Adapter),
		builder,
	)
}
