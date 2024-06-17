package kinds

import "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/outputs/kinds"

// NewAdapter creates a new adapter
func NewAdapter() kinds.Adapter {
	builder := kinds.NewBuilder()
	return createAdapter(
		builder,
	)
}
