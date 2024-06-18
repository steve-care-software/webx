package links

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/links"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/links/layers"
	json_links "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links"
)

// NewAdapter creates a new adapter
func NewAdapter() links.Adapter {
	layerAdapter := json_layers.NewAdapter()
	linkAdapter := json_links.NewAdapter()
	builder := links.NewBuilder()
	return createAdapter(
		layerAdapter.(*json_layers.Adapter),
		linkAdapter.(*json_links.Adapter),
		builder,
	)
}
