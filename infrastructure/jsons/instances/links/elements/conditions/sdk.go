package conditions

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	json_resources "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/conditions/resources"
)

// NewAdapter creates a new adapter
func NewAdapter() conditions.Adapter {
	resourceAdapter := json_resources.NewAdapter()
	builder := conditions.NewBuilder()
	return createAdapter(
		resourceAdapter.(*json_resources.Adapter),
		builder,
	)
}
