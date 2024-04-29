package constants

import "github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/constants"

// NewAdapter creates a new adapter
func NewAdapter() constants.Adapter {
	builder := constants.NewBuilder()
	return createAdapter(
		builder,
	)
}
