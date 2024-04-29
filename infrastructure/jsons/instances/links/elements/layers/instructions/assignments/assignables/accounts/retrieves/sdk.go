package retrieves

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/retrieves"
)

// NewAdapter creates a new adapter
func NewAdapter() retrieves.Adapter {
	builder := retrieves.NewBuilder()
	return createAdapter(
		builder,
	)
}
