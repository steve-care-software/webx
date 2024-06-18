package pointers

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers"
	json_conditions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/conditions"
)

// NewAdapter creates a new adapter
func NewAdapter() pointers.Adapter {
	conditionAdapter := json_conditions.NewAdapter()
	builder := pointers.NewBuilder()
	pointerBuilder := pointers.NewPointerBuilder()
	return createAdapter(
		conditionAdapter.(*json_conditions.Adapter),
		builder,
		pointerBuilder,
	)
}
