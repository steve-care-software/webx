package results

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/executions/results"
	json_interruptions "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/executions/results/interruptions"
	json_success "github.com/steve-care-software/datastencil/stencils/infrastructure/jsons/instances/executions/results/success"
)

// NewAdapter creates a new adapter
func NewAdapter() results.Adapter {
	interruptionAdapter := json_interruptions.NewAdapter()
	successAdapter := json_success.NewAdapter()
	builder := results.NewBuilder()
	return createAdapter(
		interruptionAdapter.(*json_interruptions.Adapter),
		successAdapter.(*json_success.Adapter),
		builder,
	)
}
