package cardinalities

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens/cardinalities"
)

// NewAdapter creates a new adapter
func NewAdapter() cardinalities.Adapter {
	builder := cardinalities.NewBuilder()
	return createAdapter(
		builder,
	)
}
