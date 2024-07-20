package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions"
	json_amounts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/amounts"
	json_begins "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/begins"
	json_executes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/executes"
	json_heads "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/heads"
	json_inits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/inits"
	json_retrieves "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/retrieves"
)

// NewAdapter creates a new adapter
func NewAdapter() executions.Adapter {
	amountAdapter := json_amounts.NewAdapter()
	beginAdapter := json_begins.NewAdapter()
	executeAdapter := json_executes.NewAdapter()
	headAdapter := json_heads.NewAdapter()
	initAdapter := json_inits.NewAdapter()
	retrieveAdapter := json_retrieves.NewAdapter()
	builder := executions.NewBuilder()
	contentBuilder := executions.NewContentBuilder()
	return createAdapter(
		amountAdapter.(*json_amounts.Adapter),
		beginAdapter.(*json_begins.Adapter),
		executeAdapter.(*json_executes.Adapter),
		headAdapter.(*json_heads.Adapter),
		initAdapter.(*json_inits.Adapter),
		retrieveAdapter.(*json_retrieves.Adapter),
		builder,
		contentBuilder,
	)
}
