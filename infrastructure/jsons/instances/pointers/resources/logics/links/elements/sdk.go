package elements

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements"
	json_conditions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/links/elements/conditions"
)

// NewAdapter creates a new adapter
func NewAdapter() elements.Adapter {
	conditionAdapter := json_conditions.NewAdapter()
	elementBuilder := elements.NewElementBuilder()
	builder := elements.NewBuilder()
	return createAdapter(
		conditionAdapter.(*json_conditions.Adapter),
		elementBuilder,
		builder,
	)
}
