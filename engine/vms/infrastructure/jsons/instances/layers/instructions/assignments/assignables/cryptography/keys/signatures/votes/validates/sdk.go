package validates

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
)

// NewAdapter creates an adapter
func NewAdapter() validates.Adapter {
	builder := validates.NewBuilder()
	return createAdapter(
		builder,
	)
}
