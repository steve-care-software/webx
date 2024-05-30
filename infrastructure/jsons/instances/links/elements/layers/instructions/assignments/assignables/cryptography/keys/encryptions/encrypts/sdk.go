package encrypts

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
)

// NewAdapter creates a new adapter
func NewAdapter() encrypts.Adapter {
	builder := encrypts.NewBuilder()
	return createAdapter(
		builder,
	)
}
