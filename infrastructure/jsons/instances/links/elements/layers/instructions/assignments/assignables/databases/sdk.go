package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases"
	json_repositories "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/databases/repositories"
	json_services "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/databases/services"
)

// NewAdapter creates a new adapter
func NewAdapter() databases.Adapter {
	repositoryAdapter := json_repositories.NewAdapter()
	serviceAdapter := json_services.NewAdapter()
	builder := databases.NewBuilder()
	return createAdapter(
		repositoryAdapter.(*json_repositories.Adapter),
		serviceAdapter.(*json_services.Adapter),
		builder,
	)
}
