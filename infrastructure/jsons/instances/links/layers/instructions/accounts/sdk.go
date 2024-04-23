package accounts

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts"
	json_inserts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/accounts/inserts"
	json_updates "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/accounts/updates"
)

// NewAdapter creates a new adapter
func NewAdapter() accounts.Adapter {
	insertAdapter := json_inserts.NewAdapter()
	updateAdapter := json_updates.NewAdapter()
	builder := accounts.NewBuilder()
	return createAdapter(
		insertAdapter.(*json_inserts.Adapter),
		updateAdapter.(*json_updates.Adapter),
		builder,
	)
}
