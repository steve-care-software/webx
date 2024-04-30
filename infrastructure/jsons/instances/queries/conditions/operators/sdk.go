package operators

import (
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators"
	json_integers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/queries/conditions/operators/integers"
	json_relationals "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/queries/conditions/operators/relationals"
)

// NewAdapter creates a new adapter
func NewAdapter() operators.Adapter {
	integerAdapter := json_integers.NewAdapter()
	relationalAdapter := json_relationals.NewAdapter()
	builder := operators.NewBuilder()
	return createAdapter(
		integerAdapter.(*json_integers.Adapter),
		relationalAdapter.(*json_relationals.Adapter),
		builder,
	)
}
