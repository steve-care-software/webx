package operators

import "github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"

// NewAdapter creates a new adapter
func NewAdapter() operators.Adapter {
	builder := operators.NewBuilder()
	return createAdapter(
		builder,
	)
}
