package outputs

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/success/outputs"
)

// NewAdapter creates a new adapter
func NewAdapter() outputs.Adapter {
	builder := outputs.NewBuilder()
	return createAdapter(
		builder,
	)
}
