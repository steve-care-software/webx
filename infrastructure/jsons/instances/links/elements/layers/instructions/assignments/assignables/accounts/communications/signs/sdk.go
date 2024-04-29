package signs

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/signs"
)

// NewAdapter creates a new adapter
func NewAdapter() signs.Adapter {
	builder := signs.NewBuilder()
	return createAdapter(
		builder,
	)
}
