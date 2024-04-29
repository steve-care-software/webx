package origins

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/origins"
	json_operators "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/origins/operators"
	json_resources "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/origins/resources"
)

// NewAdapter creates a new adapter
func NewAdapter() origins.Adapter {
	resourceAdapter := json_resources.NewAdapter()
	operatorAdapter := json_operators.NewAdapter()
	valueBuilder := origins.NewValueBuilder()
	builder := origins.NewBuilder()
	return createAdapter(
		resourceAdapter.(*json_resources.Adapter),
		operatorAdapter.(*json_operators.Adapter),
		valueBuilder,
		builder,
	)
}
