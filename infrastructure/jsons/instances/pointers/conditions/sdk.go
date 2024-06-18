package conditions

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions"
	json_operators "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/conditions/operators"
)

// NewAdapter creates a new adapter
func NewAdapter() conditions.Adapter {
	operatorAdapter := json_operators.NewAdapter()
	builder := conditions.NewBuilder()
	resourceBuilder := conditions.NewResourceBuilder()
	comparisonsBuilder := conditions.NewComparisonsBuilder()
	comparisonBuilder := conditions.NewComparisonBuilder()
	return createAdapter(
		operatorAdapter.(*json_operators.Adapter),
		builder,
		resourceBuilder,
		comparisonsBuilder,
		comparisonBuilder,
	)
}
