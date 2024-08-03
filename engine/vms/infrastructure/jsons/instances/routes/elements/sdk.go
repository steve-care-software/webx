package elements

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
)

// NewAdapter creates a new adapter
func NewAdapter() elements.Adapter {
	hashAdapter := hash.NewAdapter()
	builder := elements.NewBuilder()
	elementBuilder := elements.NewElementBuilder()
	return createAdapter(
		hashAdapter,
		builder,
		elementBuilder,
	)
}
