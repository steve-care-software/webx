package creates

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
)

// NewAdapter creates a new adapter
func NewAdapter() creates.Adapter {
	builder := creates.NewBuilder()
	return createAdapter(
		builder,
	)
}
