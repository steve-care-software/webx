package executions

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions"
	json_executes "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/executes"
	json_inits "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/inits"
	json_retrieves "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/retrieves"
)

// NewAdapter creates a new adapter
func NewAdapter() executions.Adapter {
	executeAdapter := json_executes.NewAdapter()
	initAdapter := json_inits.NewAdapter()
	retrieveAdapter := json_retrieves.NewAdapter()
	builder := executions.NewBuilder()
	contentBuilder := executions.NewContentBuilder()
	return createAdapter(
		executeAdapter.(*json_executes.Adapter),
		initAdapter.(*json_inits.Adapter),
		retrieveAdapter.(*json_retrieves.Adapter),
		builder,
		contentBuilder,
	)
}
