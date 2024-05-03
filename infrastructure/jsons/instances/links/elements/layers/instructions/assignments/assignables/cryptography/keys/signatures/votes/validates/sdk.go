package validates

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
)

// NewAdapter creates an adapter
func NewAdapter() validates.Adapter {
	builder := validates.NewBuilder()
	return createAdapter(
		builder,
	)
}
