package creates

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
)

// NewAdapter creates a new adapter
func NewAdapter() creates.Adapter {
	builder := creates.NewBuilder()
	return createAdapter(
		builder,
	)
}
