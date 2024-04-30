package integers

import "github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/integers"

// NewAdapter creates a new adapter
func NewAdapter() integers.Adapter {
	builder := integers.NewBuilder()
	return createAdapter(
		builder,
	)
}
