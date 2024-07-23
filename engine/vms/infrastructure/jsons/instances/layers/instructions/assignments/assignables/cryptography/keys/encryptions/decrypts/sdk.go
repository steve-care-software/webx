package decrypts

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
)

// NewAdapter creates a new adapter
func NewAdapter() decrypts.Adapter {
	builder := decrypts.NewBuilder()
	return createAdapter(
		builder,
	)
}
