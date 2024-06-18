package links

import (
	"github.com/steve-care-software/datastencil/domain/instances/links"
	json_elements "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements"
	json_references "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/references"
)

// NewAdapter creates a new adapter
func NewAdapter() links.Adapter {
	elementAdapter := json_elements.NewAdapter()
	referenceAdapter := json_references.NewAdapter()
	builder := links.NewBuilder()
	return createAdapter(
		elementAdapter.(*json_elements.Adapter),
		referenceAdapter.(*json_references.Adapter),
		builder,
	)
}
