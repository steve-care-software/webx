package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases"
	json_actions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/databases/actions"
	json_commits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/databases/commits"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/databases/databases"
	json_deletes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/databases/deletes"
	json_modifications "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/databases/modifications"
	json_retrieves "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/databases/retrieves"
)

// NewAdapter creates a new adapter
func NewAdapter() databases.Adapter {
	actionAdapter := json_actions.NewAdapter()
	commitAdapter := json_commits.NewAdapter()
	databaseAdapter := json_databases.NewAdapter()
	deleteAdapter := json_deletes.NewAdapter()
	modificationAdapter := json_modifications.NewAdapter()
	retrieveAdapter := json_retrieves.NewAdapter()
	builder := databases.NewBuilder()
	return createAdapter(
		actionAdapter.(*json_actions.Adapter),
		commitAdapter.(*json_commits.Adapter),
		databaseAdapter.(*json_databases.Adapter),
		deleteAdapter.(*json_deletes.Adapter),
		modificationAdapter.(*json_modifications.Adapter),
		retrieveAdapter.(*json_retrieves.Adapter),
		builder,
	)
}
