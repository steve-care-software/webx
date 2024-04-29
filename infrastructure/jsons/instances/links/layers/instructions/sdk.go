package instructions

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions"
	json_accounts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/accounts"
	json_assignments "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/databases"
)

// NewAdapter creates a new adapter
func NewAdapter() instructions.Adapter {
	accountAdapter := json_accounts.NewAdapter()
	assignmnetAdapter := json_assignments.NewAdapter()
	databaseAdapter := json_databases.NewAdapter()
	conditionBuilder := instructions.NewConditionBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	builder := instructions.NewBuilder()
	return createAdapter(
		accountAdapter.(*json_accounts.Adapter),
		assignmnetAdapter.(*json_assignments.Adapter),
		databaseAdapter.(*json_databases.Adapter),
		conditionBuilder,
		instructionBuilder,
		builder,
	)
}
