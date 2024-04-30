package instructions

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions"
	json_assignments "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments"
)

// NewAdapter creates a new adapter
func NewAdapter() instructions.Adapter {
	assignmnetAdapter := json_assignments.NewAdapter()
	conditionBuilder := instructions.NewConditionBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	builder := instructions.NewBuilder()
	return createAdapter(
		assignmnetAdapter.(*json_assignments.Adapter),
		conditionBuilder,
		instructionBuilder,
		builder,
	)
}
