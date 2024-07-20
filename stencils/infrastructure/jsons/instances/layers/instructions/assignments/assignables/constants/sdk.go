package constants

import "github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/constants"

// NewAdapter creates a new adapter
func NewAdapter() constants.Adapter {
	builder := constants.NewBuilder()
	constantBuilder := constants.NewConstantBuilder()
	return createAdapter(
		builder,
		constantBuilder,
	)
}
