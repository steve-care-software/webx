package instructions

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions"
	json_assignments "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/databases"
	json_lists "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/lists"
)

// NewAdapter creates a new adapter
func NewAdapter() instructions.Adapter {
	assignmnetAdapter := json_assignments.NewAdapter()
	databaseAdapter := json_databases.NewAdapter()
	listAdapter := json_lists.NewAdapter()
	loopBuilder := instructions.NewLoopBuuilder()
	conditionBuilder := instructions.NewConditionBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	builder := instructions.NewBuilder()
	return createAdapter(
		assignmnetAdapter.(*json_assignments.Adapter),
		databaseAdapter.(*json_databases.Adapter),
		listAdapter.(*json_lists.Adapter),
		loopBuilder,
		conditionBuilder,
		instructionBuilder,
		builder,
	)
}
