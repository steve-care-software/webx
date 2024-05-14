package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	json_commits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits"
)

// NewAdapter creates a new adapter
func NewAdapter() databases.Adapter {
	commitAdapter := json_commits.NewAdapter()
	builder := databases.NewBuilder()
	return createAdapter(
		commitAdapter.(*json_commits.Adapter),
		builder,
	)
}
