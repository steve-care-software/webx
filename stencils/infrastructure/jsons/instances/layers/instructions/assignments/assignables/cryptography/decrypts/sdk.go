package decrypts

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
)

// NewAdapter creates a new adapter
func NewAdapter() decrypts.Adapter {
	builder := decrypts.NewBuilder()
	return createAdapter(
		builder,
	)
}
