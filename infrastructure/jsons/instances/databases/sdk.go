package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	json_commits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits"
	json_heads "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/heads"
)

// NewAdapter creates a new adapter
func NewAdapter() databases.Adapter {
	commitAdapter := json_commits.NewAdapter()
	headAdapter := json_heads.NewAdapter()
	builder := databases.NewBuilder()
	return createAdapter(
		commitAdapter.(*json_commits.Adapter),
		headAdapter.(*json_heads.Adapter),
		builder,
	)
}
