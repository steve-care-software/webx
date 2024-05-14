package modifications

import (
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	json_deletes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits/actions/modifications/deletes"
)

// NewAdapter creates a new adapter
func NewAdapter() modifications.Adapter {
	deleteAdapter := json_deletes.NewAdapter()
	modificationBuilder := modifications.NewModificationBuilder()
	builder := modifications.NewBuilder()
	return createAdapter(
		deleteAdapter.(*json_deletes.Adapter),
		modificationBuilder,
		builder,
	)
}
