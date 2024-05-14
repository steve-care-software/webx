package commits

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	json_actions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits/actions"
)

// NewAdapter creates a new adapter
func NewAdapter() commits.Adapter {
	actionAdapter := json_actions.NewAdapter()
	builder := commits.NewBuilder()
	hashAdapter := hash.NewAdapter()
	return createAdapter(
		actionAdapter.(*json_actions.Adapter),
		builder,
		hashAdapter,
	)
}
