package instructions

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions"
	json_assignments "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments"
	json_executions "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/executions"
	json_lists "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/lists"
)

// NewAdapter creates a new adapter
func NewAdapter() instructions.Adapter {
	executionAdapter := json_executions.NewAdapter()
	assignmnetAdapter := json_assignments.NewAdapter()
	listAdapter := json_lists.NewAdapter()
	loopBuilder := instructions.NewLoopBuuilder()
	conditionBuilder := instructions.NewConditionBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	builder := instructions.NewBuilder()
	return createAdapter(
		executionAdapter.(*json_executions.Adapter),
		assignmnetAdapter.(*json_assignments.Adapter),
		listAdapter.(*json_lists.Adapter),
		loopBuilder,
		conditionBuilder,
		instructionBuilder,
		builder,
	)
}
