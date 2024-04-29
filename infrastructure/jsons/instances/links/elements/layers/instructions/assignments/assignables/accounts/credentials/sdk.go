package credentials

import "github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/credentials"

// NewAdapter creates a new adapter
func NewAdapter() credentials.Adapter {
	builder := credentials.NewBuilder()
	return createAdapter(
		builder,
	)
}
