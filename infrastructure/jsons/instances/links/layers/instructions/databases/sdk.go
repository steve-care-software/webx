package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases"
	json_deletes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/databases/deletes"
	json_inserts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/databases/inserts"
	json_reverts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/databases/reverts"
)

// NewAdapter creates a new adapter
func NewAdapter() databases.Adapter {
	insertAdapter := json_inserts.NewAdapter()
	deleteAdapter := json_deletes.NewAdapter()
	revertAdapter := json_reverts.NewAdapter()
	builder := databases.NewBuilder()
	return createAdapter(
		insertAdapter.(*json_inserts.Adapter),
		deleteAdapter.(*json_deletes.Adapter),
		revertAdapter.(*json_reverts.Adapter),
		builder,
	)
}
