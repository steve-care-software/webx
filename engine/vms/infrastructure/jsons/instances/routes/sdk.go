package routes

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes"
	json_omissions "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/omissions"
	json_tokens "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/tokens"
)

// NewAdapter creates a new adapter
func NewAdapter() routes.Adapter {
	omissionAdapter := json_omissions.NewAdapter()
	tokenAdapter := json_tokens.NewAdapter()
	builder := routes.NewBuilder()
	hashAdapter := hash.NewAdapter()
	return createAdapter(
		omissionAdapter.(*json_omissions.Adapter),
		tokenAdapter.(*json_tokens.Adapter),
		builder,
		hashAdapter,
	)
}
