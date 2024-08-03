package contexts

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/contexts"
)

// NewAdapter creates a new adapter
func NewAdapter() contexts.Adapter {
	builder := contexts.NewBuilder()
	hashAdapter := hash.NewAdapter()
	return createAdapter(
		builder,
		hashAdapter,
	)
}
