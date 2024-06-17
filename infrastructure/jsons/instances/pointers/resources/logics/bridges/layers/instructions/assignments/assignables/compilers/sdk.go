package compilers

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/compilers"
)

// NewAdapter creates a new adapter
func NewAdapter() compilers.Adapter {
	builder := compilers.NewBuilder()
	return createAdapter(
		builder,
	)
}
