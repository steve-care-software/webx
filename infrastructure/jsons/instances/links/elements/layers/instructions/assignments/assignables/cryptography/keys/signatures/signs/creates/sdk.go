package creates

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
)

// NewAdapter creates a new adapter
func NewAdapter() creates.Adapter {
	builder := creates.NewBuilder()
	return createAdapter(
		builder,
	)
}
