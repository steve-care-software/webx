package instructions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions"
	json_assignments "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments"
	json_lists "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/lists"
)

// NewAdapter creates a new adapter
func NewAdapter() instructions.Adapter {
	assignmnetAdapter := json_assignments.NewAdapter()
	listAdapter := json_lists.NewAdapter()
	loopBuilder := instructions.NewLoopBuuilder()
	conditionBuilder := instructions.NewConditionBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	builder := instructions.NewBuilder()
	return createAdapter(
		assignmnetAdapter.(*json_assignments.Adapter),
		listAdapter.(*json_lists.Adapter),
		loopBuilder,
		conditionBuilder,
		instructionBuilder,
		builder,
	)
}
