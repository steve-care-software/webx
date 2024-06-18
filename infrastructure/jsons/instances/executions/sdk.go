package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases"
	json_links "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/links"
)

// NewAdapter creates a new adapter
func NewAdapter() executions.Adapter {
	databaseAdapter := json_databases.NewAdapter()
	linkAdapter := json_links.NewAdapter()
	builder := executions.NewBuilder()
	executionBuilder := executions.NewExecutionBuilder()
	return createAdapter(
		databaseAdapter.(*json_databases.Adapter),
		linkAdapter.(*json_links.Adapter),
		builder,
		executionBuilder,
	)
}
