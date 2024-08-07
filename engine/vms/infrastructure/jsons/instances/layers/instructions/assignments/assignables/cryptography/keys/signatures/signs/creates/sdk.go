package creates

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
)

// NewAdapter creates a new adapter
func NewAdapter() creates.Adapter {
	builder := creates.NewBuilder()
	return createAdapter(
		builder,
	)
}
