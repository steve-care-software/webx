package modifications

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/modifications"
)

// NewAdapter creates a new adapter
func NewAdapter() modifications.Adapter {
	builder := modifications.NewBuilder()
	return createAdapter(
		builder,
	)
}
