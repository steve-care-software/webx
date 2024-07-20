package outputs

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/success/outputs"
)

// NewAdapter creates a new adapter
func NewAdapter() outputs.Adapter {
	builder := outputs.NewBuilder()
	return createAdapter(
		builder,
	)
}
