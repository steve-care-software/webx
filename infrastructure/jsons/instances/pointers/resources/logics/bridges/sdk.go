package bridges

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers"
)

// NewAdapter creates a new adapter
func NewAdapter() bridges.Adapter {
	layerAdapter := json_layers.NewAdapter()
	builder := bridges.NewBuiler()
	bridgeBuilder := bridges.NewBridgeBuilder()
	return createAdapter(
		layerAdapter.(*json_layers.Adapter),
		builder,
		bridgeBuilder,
	)
}
