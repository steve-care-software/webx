package bytes

import (
	infra_bytes "github.com/steve-care-software/webx/engine/databases/bytes/infrastructure/bytes"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/pointers"
)

// NewPointerAdapter creates a new pointer adapter
func NewPointerAdapter() pointers.Adapter {
	hashAdater := hash.NewAdapter()
	delimiterAdapter := infra_bytes.NewDelimiterAdapter()
	builder := pointers.NewBuilder()
	pointerBuilder := pointers.NewPointerBuilder()
	return createPointerAdapter(
		hashAdater,
		delimiterAdapter,
		builder,
		pointerBuilder,
	)
}
