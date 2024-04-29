package repositories

import "github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/repositories"

// NewAdapter creates a new adapter
func NewAdapter() repositories.Adapter {
	builder := repositories.NewBuilder()
	return createAdapter(
		builder,
	)
}
