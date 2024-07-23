package inserts

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/lists/inserts"
)

// NewAdapter creates a new adapter
func NewAdapter() inserts.Adapter {
	builder := inserts.NewBuilder()
	return createAdapter(
		builder,
	)
}
