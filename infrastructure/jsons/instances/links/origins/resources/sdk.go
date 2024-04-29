package resources

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/resources"
)

// NewAdapter creates a new adapter
func NewAdapter() resources.Adapter {
	hashAdapter := hash.NewAdapter()
	builder := resources.NewBuilder()
	return createAdapter(
		hashAdapter,
		builder,
	)
}
