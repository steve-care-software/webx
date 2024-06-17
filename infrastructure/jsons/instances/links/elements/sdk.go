package elements

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements"
	json_conditions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/conditions"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers"
)

// NewAdapter creates a new adapter
func NewAdapter() elements.Adapter {
	layerAdapter := json_layers.NewAdapter()
	conditionAdapter := json_conditions.NewAdapter()
	elementBuilder := elements.NewElementBuilder()
	builder := elements.NewBuilder()
	return createAdapter(
		layerAdapter.(*json_layers.Adapter),
		conditionAdapter.(*json_conditions.Adapter),
		elementBuilder,
		builder,
	)
}
