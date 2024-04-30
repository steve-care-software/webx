package links

import (
	"github.com/steve-care-software/datastencil/domain/instances/links"
	json_elements "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements"
	json_origins "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/origins"
	json_references "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/references"
)

// NewAdapter creates a new adapter
func NewAdapter() links.Adapter {
	elementAdapter := json_elements.NewAdapter()
	originAdapter := json_origins.NewAdapter()
	referenceAdapter := json_references.NewAdapter()
	linkBuilder := links.NewLinkBuilder()
	builder := links.NewBuilder()
	return createAdapter(
		elementAdapter.(*json_elements.Adapter),
		originAdapter.(*json_origins.Adapter),
		referenceAdapter.(*json_references.Adapter),
		linkBuilder,
		builder,
	)
}
