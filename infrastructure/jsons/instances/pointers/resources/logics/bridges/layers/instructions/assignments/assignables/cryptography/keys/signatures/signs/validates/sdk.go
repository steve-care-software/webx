package validates

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
)

// NewAdapter creates a new adapter
func NewAdapter() validates.Adapter {
	builder := validates.NewBuilder()
	return createAdapter(
		builder,
	)
}
