package encrypts

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
)

// NewAdapter creates a new adapter
func NewAdapter() encrypts.Adapter {
	builder := encrypts.NewBuilder()
	return createAdapter(
		builder,
	)
}
