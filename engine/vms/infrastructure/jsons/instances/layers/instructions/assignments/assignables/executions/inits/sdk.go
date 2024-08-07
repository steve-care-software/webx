package inits

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/executions/inits"
)

// NewAdapter creates a new adapter
func NewAdapter() inits.Adapter {
	builder := inits.NewBuilder()
	return createAdapter(
		builder,
	)
}
