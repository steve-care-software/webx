package bytes

import (
	infra_bytes "github.com/steve-care-software/webx/engine/bytes/infrastructure/bytes"
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/hashes/domain/pointers"
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
