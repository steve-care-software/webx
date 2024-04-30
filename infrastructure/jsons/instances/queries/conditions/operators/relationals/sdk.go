package relationals

import "github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/relationals"

// NewAdapter creates a new adapter
func NewAdapter() relationals.Adapter {
	builder := relationals.NewBuilder()
	return createAdapter(
		builder,
	)
}
