package updates

import (
	json_criterias "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/accounts/updates/criterias"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/accounts/updates"
)

// NewAdapter creates a new adapter
func NewAdapter() updates.Adapter {
	criteriaAdapter := json_criterias.NewAdapter()
	builder := updates.NewBuilder()
	return createAdapter(
		criteriaAdapter.(*json_criterias.Adapter),
		builder,
	)
}
