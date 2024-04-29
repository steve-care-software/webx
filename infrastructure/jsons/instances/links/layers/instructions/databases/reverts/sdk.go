package reverts

import "github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/reverts"

// NewAdapter creates a new adapter
func NewAdapter() reverts.Adapter {
	builder := reverts.NewBuilder()
	return createAdapter(
		builder,
	)
}
