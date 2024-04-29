package outputs

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs"
	json_kinds "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/outputs/kinds"
)

// NewAdapter creates a new adapter
func NewAdapter() outputs.Adapter {
	kindAdapter := json_kinds.NewAdapter()
	builder := outputs.NewBuilder()
	return createAdapter(
		kindAdapter.(*json_kinds.Adapter),
		builder,
	)
}
