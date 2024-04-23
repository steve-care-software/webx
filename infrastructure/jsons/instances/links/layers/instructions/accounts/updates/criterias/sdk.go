package criterias

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts/updates/criterias"
)

// NewAdapter creates a new adapter
func NewAdapter() criterias.Adapter {
	builder := criterias.NewBuilder()
	return createAdapter(
		builder,
	)
}
