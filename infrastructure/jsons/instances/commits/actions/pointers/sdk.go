package pointers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/pointers"
)

// NewAdapter creates a new adapter
func NewAdapter() pointers.Adapter {
	hashAdapter := hash.NewAdapter()
	builder := pointers.NewBuilder()
	return createAdapter(
		hashAdapter,
		builder,
	)
}
