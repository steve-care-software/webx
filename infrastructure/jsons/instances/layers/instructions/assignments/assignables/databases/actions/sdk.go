package actions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/actions"
)

// NewAdapter creates a new adapter
func NewAdapter() actions.Adapter {
	builder := actions.NewBuilder()
	return createAdapter(
		builder,
	)
}
