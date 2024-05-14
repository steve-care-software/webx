package actions

import (
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	json_modifications "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits/actions/modifications"
)

// NewAdapter creates a new adapter
func NewAdapter() actions.Adapter {
	modificationAdapter := json_modifications.NewAdapter()
	actionBuilder := actions.NewActionBuilder()
	builder := actions.NewBuilder()
	return createAdapter(
		modificationAdapter.(*json_modifications.Adapter),
		actionBuilder,
		builder,
	)
}
